package dao

import (
	l "crossing-api/libs/log"
	m "crossing-api/models"
)

// GetCounties fetches the metadata's counties
func GetCounties(counties *[]m.County, state string) (err error) {
	l.Info("Fetching US counties for state:", state)
	geographicInfo := metadataClient.Child(geographicInfoBucket)
	if err := geographicInfo.Child(countiesBucket).Child(state).Get(ctx, &counties); err != nil {
		l.Error("Error reading states:", err)
		return err
	}
	return nil
}
