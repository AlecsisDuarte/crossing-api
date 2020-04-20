package routes

import (
	"github.com/AlecsisDuarte/crossing-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", controllers.V1Index)
		v1.GET("/ports", controllers.V1GetPorts)
		v1.GET("/port/:portNumber", controllers.V1GetPort)
	}
	r.GET("/", controllers.Index)
	return r
}
