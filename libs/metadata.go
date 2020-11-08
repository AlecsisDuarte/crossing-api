package libs

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"crossing-api/libs/cache"
	"crossing-api/libs/log"
	"crossing-api/models"
)

const (
	metadataCacheKey = "METADATA_CACHED"
)

// GetMetadataJSON reads the metadata json found within the DB folder and parses it to it's model
func GetMetadataJSON() *models.Metadata {
	cachedMetadata := getCachedMetadata()
	if cachedMetadata != nil {
		log.Info("Returning cached metadata")
		return cachedMetadata
	}

	metadata := readLocalMetadata()
	log.Info("Successfully read the Metadata.json")
	cacheMetada(&metadata)
	return &metadata
}

// UpdateMetadata updates the cached metadata using the local JSON file
func UpdateMetadata() {
	metadata := readLocalMetadata()
	cacheMetada(&metadata)
}

func readLocalMetadata() models.Metadata {
	metadataJSON, err := os.Open("database/metadata.json")
	if err != nil {
		log.Fatal("Error reading database default metadata.json", err)
	}

	defer metadataJSON.Close()

	metadataByteValue, _ := ioutil.ReadAll(metadataJSON)

	var metadata models.Metadata
	json.Unmarshal([]byte(metadataByteValue), &metadata)
	return metadata
}

func cacheMetada(metadata *models.Metadata) {
	cache.Put(metadataCacheKey, metadata)
}

func getCachedMetadata() (metadata *models.Metadata) {
	res, found := cache.Get(metadataCacheKey)
	if !found {
		log.Info("There is no metadata cached")
		return nil
	}

	log.Info("Metadata cached")
	return res.(*models.Metadata)
}
