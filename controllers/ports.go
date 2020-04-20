package controllers

import (
	"github.com/AlecsisDuarte/crossing-api/models"
	"github.com/AlecsisDuarte/crossing-api/utils"
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
