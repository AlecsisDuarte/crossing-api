// Package utils stores all the utilities used in the project
package utils

import (
	"os"
	"time"

	"crossing-api/libs/log"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const productionKey = "PRODUCTION"
const databaseURLKey = "DATABASE_URL"
const serviceAccountKeyPathKey = "SERVICE_ACCOUNT_KEY_PATH"
const portKey = "PORT"
const defaultCacheExpiration = "DEFAULT_CACHE_DURATION_MINUTES"
const defaultCacheCleanupInterval = "DEFAULT_CACHE_CLEANUP_INTERVAL_MINUTES"

func init() {
	log.Info("Initializing ENV values")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading ENVIRONMENT", err)
	}
}

// GetEnvString returns the ENV string value or calls os.Exit() if not found
func GetEnvString(key string) string {
	v := os.Getenv(key)
	if IsEmpty(&v) {
		log.Fatal("Couldn't find the ENVIROMENT key=%s", errors.New("Unable to find the ENVIRONMENT key"), key)
	}
	return v
}

// GetEnvBool returns the ENV bool value or calls os.Exit() if not found
func GetEnvBool(key string) bool {
	s := GetEnvString(key)
	v, err := ToBool(&s)
	if err != nil {
		log.Fatal("Error while parsing ENVIRONMENT key=%s value=%s into bool", err, key, s)
		return false
	}
	return v
}

// IsProduction returns true if the application is running in release mode or false if not or not set
func IsProduction() bool {
	s := GetEnvString(productionKey)
	if IsEmpty(&s) {
		log.Warn("Production ENV key not set, defaulting into false")
		return false
	}
	v, err := ToBool(&s)
	if err != nil {
		log.Warn("Error while casting ENV Production value into bool, defaulting into false")
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

// GetPort returns the PORT to be used on the server
func GetPort() string {
	port := os.Getenv(portKey)
	if IsEmpty(&port) {
		port = "8080"
	}
	return ":" + port
}

// GetCacheExpiration returns the DEFAULT_CACHE_DURATION_MINUTES value in ENV or returns the default 5 minutes
func GetCacheExpiration() time.Duration {
	s := os.Getenv(defaultCacheExpiration)
	defaultExpiration := 5 * time.Minute

	if IsEmpty(&s) {
		log.Warn("Default cache expiration not set, using 5 minutes as default")
		return defaultExpiration
	}

	minutes, err := ToInt(&s)
	if err != nil {
		log.Warn("Default cache expiration incorrectly set, using 5 minutes as default")
		return defaultExpiration
	}
	return time.Minute * time.Duration(minutes)
}

// GetCacheCleanupInterval returns the DEFAULT_CACHE_CLEANUP_INTERVAL_MINUTES value in ENV or returns the default 10 minutes
func GetCacheCleanupInterval() time.Duration {
	s := os.Getenv(defaultCacheCleanupInterval)
	defaultExpiration := 10 * time.Minute

	if IsEmpty(&s) {
		log.Warn("Default cache cleanup interval not set. Using 10 minutes as default")
		return defaultExpiration
	}

	minutes, err := ToInt(&s)
	if err != nil {
		log.Warn("Default cache expiration incorrectly set. Using 10 minutes as default")
		return defaultExpiration
	}
	return time.Minute * time.Duration(minutes)
}
