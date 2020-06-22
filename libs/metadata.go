package libs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"crossing-api/models"
)

// GetMetadataJSON reads the metadata json found within the DB folder and parses it to it's model
func GetMetadataJSON() *models.Metadata {
	metadataJSON, err := os.Open("database/metadata.json")
	if err != nil {
		log.Fatalln("Error reading database default metadata.json")
	}

	defer metadataJSON.Close()

	metadataByteValue, _ := ioutil.ReadAll(metadataJSON)

	var metadata models.Metadata
	json.Unmarshal([]byte(metadataByteValue), &metadata)
	log.Println("Successfully read the Metadata.json")
	return &metadata
}
