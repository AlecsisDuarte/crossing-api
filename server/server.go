// Package server handles the setup process of all the APIs
package server

import (
	"crossing-api/database"
	"crossing-api/libs"
	"crossing-api/libs/cache"
	"crossing-api/libs/log"
	"crossing-api/utils"
)

// Init fetches all the server routes and runs the server on the given port
func Init() {
	log.Info("Initializing Crossing API server")

	utils.InitEnv()
	database.Init()
	cache.Init()
	libs.UpdateMetadata()

	router := SetupRouter()
	port := utils.GetPort()
	router.Run(port)
}
