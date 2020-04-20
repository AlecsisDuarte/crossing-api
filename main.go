package main

import (
	"log"

	"./src/libs"
	"./src/models"
	"./src/routes"
	"./src/utils"
)

func init() {
	utils.LoadEnvironment()
	dbRef := libs.InitDatabase()
	models.InitClients(dbRef)
}

func main() {
	// ports := libs.FetchPorts()
	// models.UpdateAllPorts(ports)

	log.Println("Starting Crossing API server")
	router := routes.SetupRouter()
	port := utils.GetPort()
	router.Run(port)
}
