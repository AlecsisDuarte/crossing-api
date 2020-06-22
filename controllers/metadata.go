package controllers

import (
	"crossing-api/models"
	"log"

	"crossing-api/libs"
	"crossing-api/utils"

	"github.com/gin-gonic/gin"
)

// V2RefreshMetadata reads local Metadata json and uploads it into our database
func V2RefreshMetadata(c *gin.Context) {
	log.Println("Refresing metadata")
	metadata := libs.GetMetadataJSON()
	if metadata == nil {
		utils.InternalError(c, "Error while reading local metadata information")
		return
	}

	exchange := libs.FetchExchangeRate()
	if exchange == nil {
		utils.InternalError(c, "Error while fetching the exchange rate")
		return
	}

	for index, country := range metadata.GeographicInfo.Countries {
		metadata.GeographicInfo.Countries[index].Exchange = exchange.Rates[country.Currency]
	}

	log.Println("Updating metadata")
	if err := models.UploadMetadata(metadata); err != nil {
		utils.InternalError(c, "Error while updating metadata information")
		return
	}
	utils.Ok(c, "Successfully updated metadata")
}

// V2GetCountries fetches the countries from the metadata bucket
func V2GetCountries(c *gin.Context) {
	log.Println("Fetching countries")
	var countries []models.Country
	if err := models.GetCountries(&countries); err != nil {
		log.Println("Error fetching metadata countries: ", err)
		utils.InternalError(c, "Error while fetching the countries")
		return
	}
	utils.Ok(c, countries)
}

// V2GetStates fetches the states from the metadata bucket
func V2GetStates(c *gin.Context) {
	country := c.Params.ByName("countryId")
	if utils.IsEmpty(&country) {
		utils.BadRequest(c, "You must specify a valid country")
		return
	}
	log.Println("Fetching states for country:", country)
	var states []models.State
	if err := models.GetStates(&states, country); err != nil {
		utils.NotFound(c, err)
		return
	}
	utils.Ok(c, states)
}

// V2GetCounties fetches the counties from the metadata bucket
func V2GetCounties(c *gin.Context) {
	state := c.Params.ByName("stateId")
	if utils.IsEmpty(&state) {
		utils.BadRequest(c, "You must specify a valid state")
		return
	}
	log.Println("Fetching counties for state:", state)
	var counties []models.County
	if err := models.GetCounties(&counties, state); err != nil {
		utils.NotFound(c, err)
		return
	}
	utils.Ok(c, counties)
}
