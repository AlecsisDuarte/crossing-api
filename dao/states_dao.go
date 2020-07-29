package dao

import (
	l "crossing-api/libs/log"
	m "crossing-api/models"
)

// GetStates fetches the metadata's states
func GetStates(states *[]m.State, country string) (err error) {
	l.Info("Fetching US surrounding states for country:", country)
	geographicInfo := metadataClient.Child(geographicInfoBucket)
	if err := geographicInfo.Child(statesBucket).Child(country).Get(ctx, &states); err != nil {
		l.Error("Error reading states: ", err)
		return err
	}
	return nil
}
