package controllers

import (
	"crossing-api/models"
	"log"

	"crossing-api/libs"
	"crossing-api/utils"

	"github.com/gin-gonic/gin"
)

// V1RefreshMetadata reads local Metadata json and uploads it into our database
func V1RefreshMetadata(c *gin.Context) {
	log.Println("Refresing metadata")
	metadata := libs.GetMetadataJSON()
	if metadata == nil {
		utils.InternalError(c, "Error while reading local metadata information")
		return
	}

	metadata.GeographicInfo.Exchange = *libs.FetchAllExchangeRates()
	for index, country := range metadata.GeographicInfo.Countries {
		metadata.GeographicInfo.Countries[index].ExchangeRate = metadata.GeographicInfo.Exchange.Rates[country.Currency]
	}

	log.Println("Updating metadata")
	if err := models.UploadMetadata(metadata); err != nil {
		utils.InternalError(c, "Error while updating metadata information")
		return
	}
	utils.Ok(c, "Successfully updated metadata")
}

// V1GetCountries fetches the countries from the metadata bucket
func V1GetCountries(c *gin.Context) {
	log.Println("Fetching countries")
	expandedStr := c.DefaultQuery("expanded", "false")
	expanded := utils.ToBoolOrDefault(&expandedStr, false)
	var countries []models.Country
	if err := models.GetCountries(&countries); err != nil {
		log.Println("Error fetching metadata countries: ", err)
		utils.InternalError(c, "Error while fetching the countries")
		return
	}
	if expanded == true {
		for countryIndex, country := range countries {
			if err := models.GetStates(&countries[countryIndex].States, country.ID); err != nil {
				log.Println("Error fetching metadata states for:", country.Name)
				utils.InternalError(c, "Error while fetching the states")
				return
			}
			for stateIndex, state := range countries[countryIndex].States {
				if err := models.GetCounties(&countries[countryIndex].States[stateIndex].Counties, state.ID); err != nil {
					log.Println("Error fetching metadata states for:", state.Name)
					utils.InternalError(c, "Error while fetching the states")
					return
				}
			}
		}
	}
	var exchange models.Exchange
	if err := models.GetExchange(&exchange); err != nil {
		log.Println("Error fetching the exchange rates")
		utils.InternalError(c, "Error while fetching the exchange rates")
		return
	}
	var geographicInfo models.GeographicInfo
	geographicInfo.Exchange = exchange
	geographicInfo.Countries = countries
	utils.Ok(c, geographicInfo)
}

// V1GetStates fetches the states from the metadata bucket
func V1GetStates(c *gin.Context) {
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

// V1GetCounties fetches the counties from the metadata bucket
func V1GetCounties(c *gin.Context) {
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
