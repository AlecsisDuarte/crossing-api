package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const PRODUCTION_KEY = "PRODUCTION"
const DATABASE_URL_KEY = "DATABASE_URL"
const SERVICE_ACCOUNT_KEY_PATH_KEY = "SERVICE_ACCOUNT_KEY_PATH"

func GetEnvString(key string) string {
	v := os.Getenv(key)
	if IsEmpty(&key) {
		log.Fatalf("Couldn't find the ENVIROMENT key=%s", key)
	}
	return v
}

func GetEnvBool(key string) bool {
	s := GetEnvString(key)
	v, err := ToBool(&s)
	if err != nil {
		log.Fatalf("Error while pasring ENVIRONMENT key=%s value=%s into bool\n", key, s)
		return false
	}
	return v
}

func IsProduction() bool {
	s := GetEnvString(PRODUCTION_KEY)
	if IsEmpty(&s) {
		return false
	}
	v, err := ToBool(&s)
	if err != nil {
		return false
	}
	return v
}

func GetServiceAccountKeyPath() string {
	return GetEnvString(SERVICE_ACCOUNT_KEY_PATH_KEY)
}

func GetDatabaseURL() string {
	return GetEnvString(DATABASE_URL_KEY)
}

func LoadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading ENVIRONMENT:", err)
	}
}
