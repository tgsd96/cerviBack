package app

import (
	"fmt"
	"path/filepath"

	"firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func InitFirebase(path string) *firebase.App {
	abspath, _ := filepath.Abs(path)
	var opt = option.WithCredentialsFile(abspath)
	App, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("Error while connecting to app : %s \n", err.Error())
	}
	return App
}
