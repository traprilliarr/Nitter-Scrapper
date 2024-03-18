package main

import (
	"fmt"
	arutek "github.com/arutek/backend-go-package"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/eyenote-corp/nitter-scrapper/lib"
	"gitlab.com/eyenote-corp/nitter-scrapper/route"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var S3_CLIENT *s3.S3

func main() {
	// THIS BLOCK IS MANDATORY
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		arutek.Logger("unable to get the current filename")
	}
	dirPath := filepath.Dir(filename)
	envPath := fmt.Sprint(dirPath, "/.env")
	if err := godotenv.Load(envPath); err != nil {
		arutek.Logger("Error loading .env file")
		panic(err)
	}
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		arutek.LoggerErr("Error loading locale time")
		panic(err)
	}
	time.Local = location

	// Core packages initialization
	port := os.Getenv("APP_PORT")
	r := gin.Default()

	// S3 Initialization
	// TODO this Initialization use global variable,
	// which makes the code less modular and bug prone
	lib.InitS3()

	// Start app
	route.MainRoute(r)
	r.Run(fmt.Sprint(":", port))
}
