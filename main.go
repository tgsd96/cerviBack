package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"firebase.google.com/go"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/julienschmidt/httprouter"
	"github.com/tgsd96/cerviBack/app"
	"github.com/tgsd96/cerviBack/models"
	"github.com/tgsd96/cerviBack/utils"
)

var App *firebase.App
var uploader *s3manager.Uploader

// intialize aws sessions
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
		Bucket: aws.String("cerbackimageuploads"),
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
	//
	// dst, err := os.Create(filename + extension)
	// defer dst.Close()
	// fmt.Printf("\nSaving file, %s ", filename)
	// if _, err := io.Copy(dst, file); err != nil {
	// 	fmt.Printf("Error occured during saving the file : %s", err.Error())
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	response.Status = "Success"
	response.Token = "bhakk ho gya kaam"
	response.ImageUID = filename
	msg, _ := json.Marshal(response)
	fmt.Fprint(w, string(msg))
}

func testAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var response models.ApiResponse
	response.Status = "Test Passed"
	response.Token = "eihjkdahfiueru2u3219319238902"
	response.ImageUID = ""
	msg, _ := json.Marshal(response)
	fmt.Fprint(w, string(msg))
}

func main() {
	App = app.InitFirebase()

	//
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		fmt.Printf("Error while connecting to aws: %s", err.Error())
	}
	uploader = s3manager.NewUploader(sess)
	router := httprouter.New()
	router.POST("/api/upload", uploadImage)
	router.GET("/api/test", testAPI)
	fmt.Println("Starting server :8080")
	http.ListenAndServe(":8080", router)
}
