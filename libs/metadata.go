package libs

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"crossing-api/libs/log"
	"crossing-api/models"
)

// GetMetadataJSON reads the metadata json found within the DB folder and parses it to it's model
func GetMetadataJSON() *models.Metadata {
	metadataJSON, err := os.Open("database/metadata.json")
	if err != nil {
		log.Fatal("Error reading database default metadata.json", err)
	}

	defer metadataJSON.Close()

	metadataByteValue, _ := ioutil.ReadAll(metadataJSON)

	var metadata models.Metadata
	json.Unmarshal([]byte(metadataByteValue), &metadata)
	log.Info("Successfully read the Metadata.json")
	return &metadata
}
