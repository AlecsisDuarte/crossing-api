package controllers

import (
	"fmt"
	"log"

	"crossing-api/libs"
	"crossing-api/models"
	"crossing-api/utils"

	"github.com/gin-gonic/gin"
)

// V1GetPorts returns a list of PortCBP in a JSON structure
func V1GetPorts(c *gin.Context) {
	var ports []models.PortCBP
	if err := models.GetAllPorts(&ports); err != nil {
		utils.NotFound(c, err)
		return
	}
	utils.Ok(c, ports)
}

// V1GetPort returns the PortCBP with the specified PortNumber
func V1GetPort(c *gin.Context) {
	portNumber := c.Params.ByName("portNumber")
	if utils.IsNotInt(&portNumber) {
		utils.BadRequest(c, "You must specify a valid port number")
		return
	}
	var port models.PortCBP
	if err := models.GetPort(&port, portNumber); err != nil {
		utils.NotFound(c, err)
		return
	}
	utils.Ok(c, port)
}

// V1RefreshPorts fetches all the CBPs port and updates the values in the database
func V1RefreshPorts(c *gin.Context) {
	log.Println("Refreshing CBP ports")
	ports := libs.FetchPorts()
	if len(*ports) == 0 {
		utils.InternalError(c, "Error while fetching CBP information")
		return
	}
	log.Println("Updating CBP ports")
	if err := models.UpdateAllPorts(ports); err != nil {
		utils.InternalError(c, "Error while updated CBP information")
		return
	}

	response := fmt.Sprintf("Successfully updated %d ports", len(*ports))
	utils.Ok(c, response)
}
