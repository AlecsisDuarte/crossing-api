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
	metadata := libs.GetMetadataJSON()
	if metadata == nil {
		utils.InternalError(c, "Error while reading local metadata information")
		return
	}
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

// V1GetStates fetches the states from the metadata bucket
func V1GetStates(c *gin.Context) {
	country := c.Params.ByName("countryId")
	if utils.IsEmpty(&country) {
		utils.BadRequest(c, "You must specify a valid country")
		return
	}
	log.Info("Fetching states for country", country)
	var states []m.State
	if err := dao.GetStates(&states, country); err != nil {
		utils.NotFound(c, err)
		return
	}

	utils.Ok(c, states)
}

// V1GetCounties fetches the counties from the metadata bucket
func V1GetCounties(c *gin.Context) {
	state := c.Params.ByName("stateId")
	if utils.IsEmpty(&state) {
		utils.BadRequest(c, "You must specify a valid state")
		return
	}
	log.Info("Fetching counties for state:", state)
	var counties []m.County
	if err := dao.GetCounties(&counties, state); err != nil {
		utils.NotFound(c, err)
		return
	}
	utils.Ok(c, counties)
}
