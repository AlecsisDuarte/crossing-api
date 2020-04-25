package models

import (
	"context"
	"fmt"
	"log"

	db "firebase.google.com/go/db"
)

const portBucket = "port"

var (
	portsClient *db.Ref
	ctx         context.Context
)

// InitClients initializes all the required database clients
func InitClients(dbRef *db.Ref) {
	portsClient = dbRef.Child(portBucket)
	ctx = context.Background()
}

// UpdateAllPorts overrides all the CBP ports stored in the database or inserts them if they do not
// exists
func UpdateAllPorts(ports *[]PortCBP) (err error) {
	log.Println("Mapping all ports to their respective PortNumber")
	portMaps := make(map[string]PortCBP)
	for i, port := range *ports {
		portMaps[port.PortNumber] = (*ports)[i]
	}

	if err := portsClient.Set(ctx, portMaps); err != nil {
		log.Fatalln("Error updating ports:", err)
		return err
	}
	log.Println("Successfully updated ports")
	return nil
}

// GetAllPorts fetches the latest status of all the CBP ports
func GetAllPorts(ports *[]PortCBP) (err error) {
	portMaps := make(map[string]*PortCBP)
	if err := portsClient.Get(ctx, &portMaps); err != nil {
		log.Fatalln("Error reading value:", err)
		return err
	}
	for _, port := range portMaps {
		*ports = append(*ports, *port)
	}
	return nil
}

// GetPort fetches the port with the specified port number
func GetPort(port *PortCBP, portNumber string) (err error) {
	if err := portsClient.Child(portNumber).Get(ctx, &port); err != nil {
		log.Fatalln("Error fetching port #:", portNumber, err)
		return err
	}
	if port == nil {
		return fmt.Errorf("Port # %s not found", portNumber)
	}
	return nil
}
