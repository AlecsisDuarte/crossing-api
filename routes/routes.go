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
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/refreshMetadata", controllers.V2RefreshMetadata)
		v2.GET("/countries", controllers.V2GetCountries)
		v2.GET("/states/:countryId", controllers.V2GetStates)
		v2.GET("/counties/:stateId", controllers.V2GetCounties)
	}
	r.GET("/", controllers.Index)
	return r
}
