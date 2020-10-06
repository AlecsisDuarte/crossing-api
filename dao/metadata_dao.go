package dao

import (
	"crossing-api/libs"
	l "crossing-api/libs/log"
	m "crossing-api/models"
)

// UploadMetadata uploads metadata information to the database
func UploadMetadata(metadata *m.Metadata) (err error) {
	l.Info("Trying to upload metadata to the database")
	if err := metadataClient.Set(ctx, metadata); err != nil {
		l.Error("Error while uploading metadata information", err)
		return err
	}
	l.Info("Successfully uploaded metadata information")
	return nil
}

// GetLocalCountries returns all the countries stored in the local JSON and if expanded it includes
// the states, counties and all the ports
func GetLocalCountries(expandedCountry *m.ExpandedCountry, expanded bool) (err error) {
	metadata := libs.GetMetadataJSON()
	countries := metadata.GeographicInfo.Countries
	expandedCountry.Countries = &countries

	for countryIndex, country := range countries {
		countries[countryIndex].Exchange = *libs.FetchExchangeRate(country.Currency)
		countries[countryIndex].Exchange.Symbol = country.Currency
		if !expanded {
			// If not expanded we won't populate States nor Counties
			continue
		}
		countries[countryIndex].States = metadata.GeographicInfo.States[country.ID]
		for stateIndex, state := range countries[countryIndex].States {
			countries[countryIndex].States[stateIndex].Counties = metadata.GeographicInfo.Counties[state.ID]
		}
	}

	if !expanded {
		// If not expanded we won't fetch all the ports
		return nil
	}

	var ports []m.PortCBP
	portsPtr := &ports
	if err := GetAllPorts(&portsPtr); err != nil {
		l.Error("Error all the ports", err)
		return err
	}
	expandedCountry.Ports = portsPtr
	return nil
}
