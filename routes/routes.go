// Package routes handles the maping of the URL paths to their specific controller
package routes

import (
	"crossing-api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter starts the github.com/gin-gonic/gin engine with all the API's route
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", controllers.V1Index)
		v1.GET("/ports", controllers.V1GetPorts)
		v1.GET("/port/:portNumber", controllers.V1GetPort)
		v1.GET("/refreshPorts", controllers.V1RefreshPorts)
	}
	r.GET("/", controllers.Index)
	return r
}
