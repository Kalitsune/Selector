package api

import (
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io"
	"log"
	"os"
)

func init() {
	/*
		init the Logger
	*/
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logStream := io.MultiWriter(file, os.Stdout)

	//init the loggers
	Logger.Info = log.New(logStream, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Warning = log.New(logStream, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Error = log.New(logStream, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	Logger.Info.Println("Logger initialized!")

	/*
		init the Google API
	*/
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		Logger.Error.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err = google.ConfigFromJSON(b, drive.DriveAppdataScope)
	if err != nil {
		Logger.Error.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	Logger.Info.Println("Google API initialized!")
}
