package controllers

import (
	"fmt"
	"log"

	"crossing-api/dao"
	"crossing-api/libs"
	m "crossing-api/models"
	"crossing-api/utils"

	"github.com/gin-gonic/gin"
)

// V1GetPorts returns a list of PortCBP in a JSON structure
func V1GetPorts(c *gin.Context) {
	var ports []m.PortCBP
	portPt := &ports
	if err := dao.GetAllPorts(&portPt); err != nil {
		utils.NotFound(c, err)
		return
	}
	utils.Ok(c, *portPt)
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
	if err := dao.UpdateAllPorts(ports); err != nil {
		utils.InternalError(c, "Error while updated CBP information")
		return
	}

	response := fmt.Sprintf("Successfully updated %d ports", len(*ports))
	utils.Ok(c, response)
}
