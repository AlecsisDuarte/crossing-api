package config

import (
	"log"
	"path/filepath"

	"github.com/AlecsisDuarte/crossing-api/utils"
)

const TEST_DATABASE_NAME = "test_crossing"
const PROD_DATABASE_NAME = "prod_crossing"

type DBConfig struct {
	DatabaseURL           string
	DatabaseName          string
	ServiceAccountKeyPath string
}

func BuildDBConfig() *DBConfig {
	relativeAccountKeyPath := utils.GetServiceAccountKeyPath()
	isProduction := utils.IsProduction()
	databaseURL := utils.GetDatabaseURL()

	serviceAccountKeyFilePath, err := filepath.Abs(relativeAccountKeyPath)
	if err != nil {
		log.Fatalln("Error couldn't find firebase service account key:", err)
	}
	databaseName := TEST_DATABASE_NAME
	if isProduction {
		databaseName = PROD_DATABASE_NAME
	}

	return &DBConfig{
		DatabaseURL:           databaseURL,
		DatabaseName:          databaseName,
		ServiceAccountKeyPath: serviceAccountKeyFilePath,
	}
}
