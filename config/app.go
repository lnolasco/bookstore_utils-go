package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// AppName - Application Name
	AppName = ""

	// ApiPort - Port number API
	AppPort = 0

	// AppEnv - Determines the "environment" the application is currently running in.
	AppEnv = ""

	// Debug - When the application is in debug mode, detailed error messages, etc.
	Debug = false

	// Url - Api Url
	Url = ""
)

// Charge App initialize the .env variables
func ChargeApp() {
	var err error
	if godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	AppPort, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		AppPort = 8080
	}

	Debug, err = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		Debug = false
	}

	AppName = os.Getenv("APP_NAME")
	AppEnv = os.Getenv("APP_ENV")
	Url = os.Getenv("APP_URL")

}
