package models

import (
	"context"
	"crossing-api/libs/log"
	"fmt"

	db "firebase.google.com/go/db"
)

const (
	portBucket           = "port"
	metadataBucket       = "metadata"
	geographicInfoBucket = "geographic_info"
	countriesBucket      = "countries"
	statesBucket         = "states"
	countiesBucket       = "counties"
	exchangeBucket       = "exchange"
)

var (
	portClient     *db.Ref
	metadataClient *db.Ref
	ctx            context.Context
)

// InitClients initializes all the required database clients
func InitClients(dbRef *db.Ref) {
	portClient = dbRef.Child(portBucket)
	metadataClient = dbRef.Child(metadataBucket)
	ctx = context.Background()
}

// UpdateAllPorts overrides all the CBP ports stored in the database or inserts them if they do not
// exists
func UpdateAllPorts(ports *[]PortCBP) (err error) {
	log.Info("Mapping all ports to their respective PortNumber")
	portMaps := make(map[string]PortCBP)
	for i, port := range *ports {
		portMaps[port.PortNumber] = (*ports)[i]
	}

	if err := portClient.Set(ctx, portMaps); err != nil {
		log.Error("Error updating ports map", err)
		return err
	}
	log.Info("Successfully updated ports")
	return nil
}

// GetAllPorts fetches the latest status of all the CBP ports
func GetAllPorts(ports *[]PortCBP) (err error) {
	portMaps := make(map[string]*PortCBP)
	if err := portClient.Get(ctx, &portMaps); err != nil {
		log.Error("Error reading value", err)
		return err
	}
	for _, port := range portMaps {
		*ports = append(*ports, *port)
	}
	return nil
}

// GetPort fetches the port with the specified port number
func GetPort(port *PortCBP, portNumber string) (err error) {
	if err := portClient.Child(portNumber).Get(ctx, &port); err != nil {
		log.Error("Error fetching port #%v", err, portNumber)
		return err
	}
	if port == nil {
		return fmt.Errorf("Port #%s not found", portNumber)
	}
	return nil
}

// GetPortsByBorder returns a list of ports with the specified border name
func GetPortsByBorder(ports *[]PortCBP, border string) (err error) {
	results, err := portClient.OrderByChild("border").EqualTo(border).GetOrdered(ctx)
	if err != nil {
		log.Error("Error querying ports by border", err)
		return err
	}
	for _, res := range results {
		var port PortCBP
		if err := res.Unmarshal(&port); err != nil {
			log.Error("Error unmarshaling the ports", err)
			return err
		}
		*ports = append(*ports, port)
	}
	return nil
}

// UploadMetadata uploads metadata information to the database
func UploadMetadata(metadata *Metadata) (err error) {
	log.Info("Trying to upload metadata to the database")
	if err := metadataClient.Set(ctx, metadata); err != nil {
		log.Error("Error while uploading metadata information", err)
		return err
	}
	log.Info("Successfully uploaded metadata information")
	return nil
}

// GetCountries fetches the metadata's countries
func GetCountries(countries *[]Country) (err error) {
	log.Info("Fetching US surrounding countries")
	if err := metadataClient.Child(geographicInfoBucket).Child(countriesBucket).Get(ctx, &countries); err != nil {
		log.Error("Error reading countries", err)
		return err
	}
	return nil
}

// GetStates fetches the metadata's states
func GetStates(states *[]State, country string) (err error) {
	log.Info("Fetching US surrounding states for country:", country)
	geographicInfo := metadataClient.Child(geographicInfoBucket)
	if err := geographicInfo.Child(statesBucket).Child(country).Get(ctx, &states); err != nil {
		log.Error("Error reading states: ", err)
		return err
	}
	return nil
}

// GetCounties fetches the metadata's counties
func GetCounties(counties *[]County, state string) (err error) {
	log.Info("Fetching US counties for state:", state)
	geographicInfo := metadataClient.Child(geographicInfoBucket)
	if err := geographicInfo.Child(countiesBucket).Child(state).Get(ctx, &counties); err != nil {
		log.Error("Error reading states:", err)
		return err
	}
	return nil
}
