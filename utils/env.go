// Package utils stores all the utilities used in the project
package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const productionKey = "PRODUCTION"
const databaseURLKey = "DATABASE_URL"
const serviceAccountKeyPathKey = "SERVICE_ACCOUNT_KEY_PATH"
const portKey = "PORT"

// GetEnvString returns the ENV string value or calls os.Exit() if not found
func GetEnvString(key string) string {
	v := os.Getenv(key)
	if IsEmpty(&key) {
		log.Fatalf("Couldn't find the ENVIROMENT key=%s", key)
	}
	return v
}

// GetEnvBool returns the ENV bool value or calls os.Exit() if not found
func GetEnvBool(key string) bool {
	s := GetEnvString(key)
	v, err := ToBool(&s)
	if err != nil {
		log.Fatalf("Error while pasring ENVIRONMENT key=%s value=%s into bool\n", key, s)
		return false
	}
	return v
}

// IsProduction returns true if the application is running in release mode or false if not or not set
func IsProduction() bool {
	s := GetEnvString(productionKey)
	if IsEmpty(&s) {
		return false
	}
	v, err := ToBool(&s)
	if err != nil {
		return false
	}
	return v
}

// GetServiceAccountKeyPath returns the path to the firebase private account key
func GetServiceAccountKeyPath() string {
	return GetEnvString(serviceAccountKeyPathKey)
}

// GetDatabaseURL returns the URL of the database we will be working with
func GetDatabaseURL() string {
	return GetEnvString(databaseURLKey)
}

// LoadEnvironment loads all .env values into runtime ENV values
func LoadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading ENVIRONMENT:", err)
	}
}

// GetPort returns the PORT to be used on the server
func GetPort() string {
	port := os.Getenv(portKey)
	if IsEmpty(&port) {
		port = "8080"
	}
	return ":" + port
}
