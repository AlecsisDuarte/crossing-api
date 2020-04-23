package models

import (
	"context"
	"log"

	db "firebase.google.com/go/db"
)

const PORT_BUCKET = "port"

var (
	portsClient *db.Ref
	ctx         context.Context
)

func InitClients(dbRef *db.Ref) {
	portsClient = dbRef.Child(PORT_BUCKET)
	ctx = context.Background()
}

func UpdateAllPorts(ports *[]PortCBP) (err error) {
	log.Println("Mapping all ports to their respective PortNumber")
	portMaps := make(map[string]*PortCBP)
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

func GetAPort(port *PortCBP, portNumber string) (err error) {
	if err := portsClient.Child(portNumber).Get(ctx, &port); err != nil {
		log.Fatalln("Error reading value:", err)
		return err
	}
	return nil
}
