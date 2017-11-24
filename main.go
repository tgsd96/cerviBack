package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"firebase.google.com/go"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/julienschmidt/httprouter"
	"github.com/tgsd96/cerviBack/app"
	"github.com/tgsd96/cerviBack/models"
	"github.com/tgsd96/cerviBack/utils"
)

// Global variables
var App *firebase.App
var db *gorm.DB
var uploader *s3manager.Uploader
var sess *session.Session

// Config files for configurations
const CONFIG_FILE_URL = "./.config/config.yaml"

// Defining constants for configs

// Config
var Config models.Config

// configs, er := yaml.ReadFile(*CONFIG_FILE_URL)

// const QUEUE_URL = "https://sqs.us-west-2.amazonaws.com/907743384002/jobQueue"

// intialize aws sessions

// API-----------
/*
	Usage : host/api/upload
	Filename: image

	send multipart file to upload the images to S3 bucket
*/
func uploadImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// authenticate the user uploading the image

	// get access token from request header
	// accessToken := r.Header.Get("Authorization")
	// accessToken = accessToken[7:]
	// fmt.Printf("\nThe received access token is: %s", accessToken)
	// authenticate the user using firebase
	// client, err := App.Auth(context.Background())
	var response models.ApiResponse
	// if err != nil {
	// 	log.Printf("\nError while creating client %s", err.Error())
	// 	response.Status = "Error creating client"
	// 	msg, _ := json.Marshal(response)
	// 	fmt.Fprint(w, string(msg))
	// 	return
	// }
	// token, er := client.VerifyIDToken(accessToken)
	// if er != nil {
	// 	log.Printf("\nError authenticating ID token: %s", er.Error())
	// 	response.Status = "Invalid_Token"
	// 	msg, _ := json.Marshal(response)
	// 	fmt.Fprint(w, string(msg))
	// 	return
	// }
	// log.Printf("\nToken authenticated: %s\n", token.UID)
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	log.Println("Saving uploads")
	err := r.ParseMultipartForm(10000)
	if err != nil {
		fmt.Printf("Error while parsing the image: %s", err.Error())
		response.Status = "Unable to parse multipart data"
		msg, _ := json.Marshal(response)
		fmt.Fprint(w, string(msg))
		return
	}
	m := r.MultipartForm
	files := m.File["image"]
	file, err := files[0].Open()
	defer file.Close()
	if err != nil {
		fmt.Println("Error occured during parsing file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("\n Receieved filename : %s", files[0].Filename)

	// upload files on aws
	filename := utils.GenerateIID()
	var extension string
	extension = filepath.Ext(files[0].Filename)
	filename = filename + extension
	log.Printf("\n Received with extension: %s", extension)
	log.Printf("\nUploading with filename: %s", filename)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Config.Aws.Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		log.Printf("\nError while uploading the file to S3: %s", err.Error())
		response.Status = "Error while connecting to S3: " + err.Error()
		msg, _ := json.Marshal(response)
		fmt.Fprint(w, string(msg))
		return
	}
	log.Printf("\nSuccessfully uploaded file to S3")

	// Add to RDS table for processing
	image := models.ImageStatus{
		ImageKey: filename,
		UserID:   "testingId",
		Status:   "INQUEUE",
	}
	err = app.AddImageToTable(db, &image)
	if err != nil {
		log.Fatalf("Error adding entry to table, error : %s", err.Error())
		response.Status = "failed_to_process_image"
		msg, _ := json.Marshal(response)
		fmt.Fprint(w, string(msg))
		return
	}

	// add message to queue
	Sqsmsg := models.SqsMessage{
		ImageKey: filename,
		UserID:   "testinguserid",
	}
	err = app.PushToSQS(sess, Config.Aws.QueueUrl, Sqsmsg)
	if err != nil {
		log.Fatalf("\nError pushing messages to queue, error : %s", err.Error())
		response.Status = "failed_sqs_push"
		msg, _ := json.Marshal(response)
		fmt.Fprint(w, string(msg))
		return
	}

	response.Status = "Success"
	response.Token = "bhakk ho gya kaam"
	response.ImageUID = filename
	msg, _ := json.Marshal(response)
	fmt.Fprint(w, string(msg))
}

// API-----------
/*
	Method : POST
	Usage : /api/status/{image key}
	Header:
		Authorisation:
			- Bearer TOKEN
*/

func StatusAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	image_key := ps.ByName("image_key")
	log.Printf("\nReceived request for : %s", image_key)
	var response models.Message
	// TODO - add authentication logic

	// check whether the image exists in the table or not
	var row models.ImageStatus

	// only for testing purpose :D
	// db, _ = app.ConnectToRDS("mysql", "localhost:3306", "root", "", "cerback")
	err := db.Where("image_key=?", image_key).First(&row).Error
	if err != nil {
		log.Printf("\n Error in finding image key: %s", err.Error())
		response.Action = "image_not_found"
		w.WriteHeader(http.StatusNotFound)
	} else {
		response.Action = "image_found"
		var result models.Result
		result.ImageUID = image_key
		result.Status = row.Status
		result.Type1 = row.Type1
		result.Type2 = row.Type2
		result.Type3 = row.Type3
		response.Data = result
	}
	msg, _ := json.Marshal(response)
	fmt.Fprintf(w, string(msg))
}

func testAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var response models.ApiResponse
	response.Status = "Test Passed"
	response.Token = "eihjkdahfiueru2u3219319238902"
	response.ImageUID = ""
	msg, _ := json.Marshal(response)
	fmt.Fprint(w, string(msg))
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "<h1>You have reached the end of the galaxy</h1><h2>No web page for you!</h2><h3>Endpoints at /api/upload</h3>")
}
func main() {

	// load configurations
	configor.Load(&Config, CONFIG_FILE_URL)
	env := os.Getenv("RUN_ENV")
	var DB_CONFIG models.DBConfig
	if env == "dev" {
		DB_CONFIG = Config.Dev
	} else {
		DB_CONFIG = Config.Prod
	}

	//Connect to firebase app
	App = app.InitFirebase(DB_CONFIG.FirebaseConfig)

	var err error
	// Connect to RDS instance
	db, err = app.ConnectToRDS(DB_CONFIG.Adapter, DB_CONFIG.DbHost, DB_CONFIG.DbUser, DB_CONFIG.DbPass, DB_CONFIG.DbName)
	if err != nil {
		log.Fatalf("\nError connecting to database, check configs. Error: %s", err.Error())
		return
	}

	sess, err = session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		fmt.Printf("Error while connecting to aws: %s", err.Error())
		return
	}
	uploader = s3manager.NewUploader(sess)
	router := httprouter.New()
	router.POST("/api/upload", uploadImage)
	router.POST("/api/status/:image_key", StatusAPI)
	router.GET("/api/test", testAPI)
	router.GET("/", index)
	fmt.Println("Starting server :8080 in " + env + " mode")
	http.ListenAndServe(":8080", router)
}
