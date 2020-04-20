package main

import (
	"log"

	"github.com/AlecsisDuarte/crossing-api/libs"
	"github.com/AlecsisDuarte/crossing-api/models"
	"github.com/AlecsisDuarte/crossing-api/routes"
	"github.com/AlecsisDuarte/crossing-api/utils"
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
