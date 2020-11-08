package controllers

import (
	m "crossing-api/models"

	"crossing-api/dao"
	"crossing-api/libs"
	"crossing-api/libs/log"
	"crossing-api/utils"

	"github.com/gin-gonic/gin"
)

// V1RefreshMetadata reads local Metadata json and updates our cached information
func V1RefreshMetadata(c *gin.Context) {
	log.Info("Refresing metadata")
	libs.UpdateMetadata()
	utils.Ok(c, "Successfully updated metadata")
}

// V1GetCountries fetches the countries from the metadata bucket
func V1GetCountries(c *gin.Context) {
	log.Info("Fetching countries")
	expandedStr := c.DefaultQuery("expanded", "false")
	expanded := utils.ToBoolOrDefault(&expandedStr, false)

	var response m.ExpandedCountry
	if err := dao.GetLocalCountries(&response, expanded); err != nil {
		log.Error("Error fetching metadata countries with expanded = %v", err, expanded)
		utils.InternalError(c, "Error while fetching the countries")
		return
	}
	utils.Ok(c, response)
}
