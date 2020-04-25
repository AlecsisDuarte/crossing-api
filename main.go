package main

import (
	"log"

	"crossing-api/libs"
	"crossing-api/models"
	"crossing-api/routes"
	"crossing-api/utils"
)

func init() {
	utils.LoadEnvironment()
	dbRef := libs.InitDatabase()
	models.InitClients(dbRef)
}

func main() {
	log.Println("Starting Crossing API server")
	router := routes.SetupRouter()
	port := utils.GetPort()
	router.Run(port)
}
