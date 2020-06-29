// Package routes handles the maping of the URL paths to their specific controller
package routes

import (
	"crossing-api/controllers"
	"crossing-api/utils"
	"log"

	"github.com/gin-gonic/gin"
)

// SetupRouter starts the github.com/gin-gonic/gin engine with all the API's route
func SetupRouter() *gin.Engine {
	if utils.IsProduction() {
		log.Println("Running GIN in release mode")
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", controllers.V1Index)
		v1.GET("/ports", controllers.V1GetPorts)
		v1.GET("/port/:portNumber", controllers.V1GetPort)
		v1.GET("/refreshPorts", controllers.V1RefreshPorts)
		v1.GET("/portsByCountry/:country", controllers.V1GetPortsByCountry)
		v1.GET("/refreshMetadata", controllers.V1RefreshMetadata)
		v1.GET("/countries/*expanded", controllers.V1GetCountries)
		v1.GET("/states/:countryId", controllers.V1GetStates)
		v1.GET("/counties/:stateId", controllers.V1GetCounties)
	}
	r.GET("/", controllers.Index)
	return r
}
