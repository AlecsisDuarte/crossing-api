package config

import (
	"log"
	"path/filepath"

	"crossing-api/utils"
)

const testDatabaseName = "test_crossing"
const prodDatabaseName = "prod_crossing"

// DBConfig stores all the information required to initialize firebase database
type DBConfig struct {
	DatabaseURL           string
	DatabaseName          string
	ServiceAccountKeyPath string
}

// BuildDBConfig initilizes the database configuration
func BuildDBConfig() *DBConfig {
	relativeAccountKeyPath := utils.GetServiceAccountKeyPath()
	isProduction := utils.IsProduction()
	databaseURL := utils.GetDatabaseURL()

	serviceAccountKeyFilePath, err := filepath.Abs(relativeAccountKeyPath)
	if err != nil {
		log.Fatalln("Error couldn't find firebase service account key:", err)
	}
	databaseName := testDatabaseName
	if isProduction {
		databaseName = prodDatabaseName
	}

	return &DBConfig{
		DatabaseURL:           databaseURL,
		DatabaseName:          databaseName,
		ServiceAccountKeyPath: serviceAccountKeyFilePath,
	}
}
