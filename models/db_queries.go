package models

import (
	"context"
	"fmt"
	"log"

	db "firebase.google.com/go/db"
)

const (
	portBucket = "port"
)

var (
	portClient *db.Ref
	ctx        context.Context
)

// InitClients initializes all the required database clients
func InitClients(dbRef *db.Ref) {
	portClient = dbRef.Child(portBucket)
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

	if err := portClient.Set(ctx, portMaps); err != nil {
		log.Panicln("Error updating ports map:", err)
		return err
	}
	log.Println("Successfully updated ports")
	return nil
}

// GetAllPorts fetches the latest status of all the CBP ports
func GetAllPorts(ports *[]PortCBP) (err error) {
	portMaps := make(map[string]*PortCBP)
	if err := portClient.Get(ctx, &portMaps); err != nil {
		log.Panicln("Error reading value:", err)
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
		log.Panicln("Error fetching port #:", portNumber, err)
		return err
	}
	if port == nil {
		return fmt.Errorf("Port # %s not found", portNumber)
	}
	return nil
}

// GetPortsByBorder returns a list of ports with the specified border name
func GetPortsByBorder(ports *[]PortCBP, border string) (err error) {
	results, err := portClient.OrderByChild("border").EqualTo(border).GetOrdered(ctx)
	if err != nil {
		log.Panicln("Error querying ports by border:", err)
		return err
	}
	for _, res := range results {
		var port PortCBP
		if err := res.Unmarshal(&port); err != nil {
			log.Panicln("Error unmarshaling the ports:", err)
			return err
		}
		*ports = append(*ports, port)
	}
	return nil
}
