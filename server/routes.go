package server

import (
	"crossing-api/controllers"
	"crossing-api/libs/log"
	"crossing-api/utils"

	"github.com/gin-gonic/gin"
)

// SetupRouter starts the github.com/gin-gonic/gin engine with all the API's route
func SetupRouter() *gin.Engine {
	if utils.IsProduction() {
		log.Info("Running GIN in release mode")
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", controllers.V1Index)
		v1.GET("/ports", controllers.V1GetPorts)
		v1.GET("/refreshPorts", controllers.V1RefreshPorts)
		v1.GET("/refreshMetadata", controllers.V1RefreshMetadata)
		v1.GET("/countries/*expanded", controllers.V1GetCountries)
	}
	r.GET("/", controllers.Index)
	return r
}
