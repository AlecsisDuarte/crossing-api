package controllers

import (
	"fmt"
	"log"

	"crossing-api/libs"
	"crossing-api/models"
	"crossing-api/utils"
	"github.com/gin-gonic/gin"
)

func V1GetPorts(c *gin.Context) {
	var ports []models.PortCBP
	if err := models.GetAllPorts(&ports); err != nil {
		utils.NotFound(c, err)
	}
	utils.Ok(c, ports)
}

func V1GetPort(c *gin.Context) {
	portNumber := c.Params.ByName("portNumber")
	if utils.IsNotInt(&portNumber) {
		utils.BadRequest(c, "You must specify the port number")
	}
	var port models.PortCBP
	if err := models.GetAPort(&port, portNumber); err != nil {
		utils.NotFound(c, err)
	}
	utils.Ok(c, port)
}

func V1RefreshPorts(c *gin.Context) {
	log.Println("Refreshing CBP ports")
	ports := libs.FetchPorts()
	if len(*ports) == 0 {
		utils.InternalError(c, "Error while fetching CBP information")
	}
	log.Println("Updating CBP ports")
	if err := models.UpdateAllPorts(ports); err != nil {
		utils.InternalError(c, "Error while updated CBP information")
	}

	response := fmt.Sprintf("Successfully updated %d ports", len(*ports))
	utils.Ok(c, response)
}
