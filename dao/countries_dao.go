package dao

import (
	l "crossing-api/libs/log"
	m "crossing-api/models"
)

// GetCountries fetches the metadata's countries
func GetCountries(countries *[]m.Country) (err error) {
	l.Info("Fetching US surrounding countries")
	if err := metadataClient.Child(geographicInfoBucket).Child(countriesBucket).Get(ctx, &countries); err != nil {
		l.Error("Error reading countries", err)
		return err
	}
	return nil
}
