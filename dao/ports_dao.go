package dao

import (
	l "crossing-api/libs/log"
	m "crossing-api/models"
	"fmt"
)

// UpdateAllPorts overrides all the CBP ports stored in the database or inserts them if they do not
// exists
func UpdateAllPorts(ports *[]m.PortCBP) (err error) {
	l.Info("Mapping all ports to their respective PortNumber")
	portMaps := make(map[string]m.PortCBP)
	for i, port := range *ports {
		portMaps[port.PortNumber] = (*ports)[i]
	}

	if err := portClient.Set(ctx, portMaps); err != nil {
		l.Error("Error updating ports map", err)
		return err
	}
	l.Info("Successfully updated ports")
	return nil
}

// GetAllPorts fetches the latest status of all the CBP ports
func GetAllPorts(ports *[]m.PortCBP) (err error) {
	portMaps := make(map[string]*m.PortCBP)
	if err := portClient.Get(ctx, &portMaps); err != nil {
		l.Error("Error reading value", err)
		return err
	}
	for _, port := range portMaps {
		*ports = append(*ports, *port)
	}
	return nil
}

// GetPort fetches the port with the specified port number
func GetPort(port *m.PortCBP, portNumber string) (err error) {
	if err := portClient.Child(portNumber).Get(ctx, &port); err != nil {
		l.Error("Error fetching port #%v", err, portNumber)
		return err
	}
	if port == nil {
		return fmt.Errorf("Port #%s not found", portNumber)
	}
	return nil
}

// GetPortsByBorder returns a list of ports with the specified border name
func GetPortsByBorder(ports *[]m.PortCBP, border string) (err error) {
	results, err := portClient.OrderByChild("border").EqualTo(border).GetOrdered(ctx)
	if err != nil {
		l.Error("Error querying ports by border", err)
		return err
	}
	for _, res := range results {
		var port m.PortCBP
		if err := res.Unmarshal(&port); err != nil {
			l.Error("Error unmarshaling the ports", err)
			return err
		}
		*ports = append(*ports, port)
	}
	return nil
}
