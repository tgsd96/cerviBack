package app

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"path/filepath"

// 	"firebase.google.com/go"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3/s3manager"
// 	"github.com/jinzhu/gorm"
// 	"github.com/julienschmidt/httprouter"
// 	"github.com/tgsd96/cerback/utils"
// 	"github.com/tgsd96/cerviBack/models"
// )

// type App struct {
// 	Router   *httprouter.Router
// 	DB       *gorm.DB
// 	Sess     *session.Session
// 	Uploader *s3manager.Uploader
// 	Firebase *firebase.App
// 	Config   *models.Config
// }

// func (a *App) Initialize(adapter, host, user, pasw, name string) {
// 	var err error
// 	a.DB, err = ConnectToRDS(adapter, host, user, pasw, name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	a.Router = httprouter.New()
// 	a.Sess, err = session.NewSession(&aws.Config{
// 		Region: aws.String("us-west-2"),
// 	})
// 	if err != nil {
// 		fmt.Printf("Error while connecting to aws: %s", err.Error())
// 		return
// 	}
// 	a.Uploader = s3manager.NewUploader(a.Sess)

// }

// func (a *App) InitializeRoutes() {
// 	a.Router.POST("/api/upload", a.uploadImage)
// 	// a.Router.POST("/api/status/:image_key", StatusAPI)
// 	// a.Router.GET("/api/test", testAPI)
// 	// a.Router.GET("/", index)
// }

// // func (a *App) Run(addr String) {

// // }
// func (a *App) uploadImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

// 	// authenticate the user uploading the image

// 	// get access token from request header
// 	// accessToken := r.Header.Get("Authorization")
// 	// accessToken = accessToken[7:]
// 	// fmt.Printf("\nThe received access token is: %s", accessToken)
// 	// authenticate the user using firebase
// 	// client, err := App.Auth(context.Background())
// 	var response models.ApiResponse
// 	// if err != nil {
// 	// 	log.Printf("\nError while creating client %s", err.Error())
// 	// 	response.Status = "Error creating client"
// 	// 	msg, _ := json.Marshal(response)
// 	// 	fmt.Fprint(w, string(msg))
// 	// 	return
// 	// }
// 	// token, er := client.VerifyIDToken(accessToken)
// 	// if er != nil {
// 	// 	log.Printf("\nError authenticating ID token: %s", er.Error())
// 	// 	response.Status = "Invalid_Token"
// 	// 	msg, _ := json.Marshal(response)
// 	// 	fmt.Fprint(w, string(msg))
// 	// 	return
// 	// }
// 	// log.Printf("\nToken authenticated: %s\n", token.UID)
// 	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
// 	log.Println("Saving uploads")
// 	err := r.ParseMultipartForm(10000)
// 	if err != nil {
// 		fmt.Printf("Error while parsing the image: %s", err.Error())
// 		response.Status = "Unable to parse multipart data"
// 		msg, _ := json.Marshal(response)
// 		fmt.Fprint(w, string(msg))
// 		return
// 	}
// 	m := r.MultipartForm
// 	files := m.File["image"]
// 	file, err := files[0].Open()
// 	defer file.Close()
// 	if err != nil {
// 		fmt.Println("Error occured during parsing file")
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	log.Printf("\n Receieved filename : %s", files[0].Filename)

// 	// upload files on aws
// 	filename := utils.GenerateIID()
// 	var extension string
// 	extension = filepath.Ext(files[0].Filename)
// 	filename = filename + extension
// 	log.Printf("\n Received with extension: %s", extension)
// 	log.Printf("\nUploading with filename: %s", filename)
// 	_, err = a.Uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(Config.Aws.Bucket),
// 		Key:    aws.String(filename),
// 		Body:   file,
// 	})
// 	if err != nil {
// 		log.Printf("\nError while uploading the file to S3: %s", err.Error())
// 		response.Status = "Error while connecting to S3: " + err.Error()
// 		msg, _ := json.Marshal(response)
// 		fmt.Fprint(w, string(msg))
// 		return
// 	}
// 	log.Printf("\nSuccessfully uploaded file to S3")

// 	// Add to RDS table for processing
// 	image := models.ImageStatus{
// 		ImageKey: filename,
// 		UserID:   "testingId",
// 		Status:   "INQUEUE",
// 	}
// 	err = AddImageToTable(a.DB, &image)
// 	if err != nil {
// 		log.Fatalf("Error adding entry to table, error : %s", err.Error())
// 		response.Status = "failed_to_process_image"
// 		msg, _ := json.Marshal(response)
// 		fmt.Fprint(w, string(msg))
// 		return
// 	}

// 	// add message to queue
// 	Sqsmsg := models.SqsMessage{
// 		ImageKey: filename,
// 		UserID:   "testinguserid",
// 	}
// 	err = PushToSQS(a.Sess, Config.Aws.QueueUrl, Sqsmsg)
// 	if err != nil {
// 		log.Fatalf("\nError pushing messages to queue, error : %s", err.Error())
// 		response.Status = "failed_sqs_push"
// 		msg, _ := json.Marshal(response)
// 		fmt.Fprint(w, string(msg))
// 		return
// 	}

// 	response.Status = "Success"
// 	response.Token = "bhakk ho gya kaam"
// 	response.ImageUID = filename
// 	msg, _ := json.Marshal(response)
// 	fmt.Fprint(w, string(msg))
// }
