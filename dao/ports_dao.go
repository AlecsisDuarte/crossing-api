package dao

import (
	"crossing-api/libs"
	"crossing-api/libs/cache"
	l "crossing-api/libs/log"
	m "crossing-api/models"
	"fmt"
)

const (
	portsByBorderKeyPrefix = "CACHED_PORTS_BY_BORDER_"
	portCachedKeyPrefix    = "CACHED_PORT_"
	allPortsCachedKey      = "ALL_CACHED_PORTS"
	allPortsMapedCachedKey = "ALL_CACHED_PORTS_MAPPED"
)

// UpdateAllPorts overrides all the CBP ports cached
func UpdateAllPorts(ports *[]m.PortCBP) (err error) {
	cacheAllPorts(ports)
	return nil
}

// GetAllPorts fetches the latest status of all the CBP ports
func GetAllPorts(ports **[]m.PortCBP) (err error) {
	cachedPorts := getAllCachedPorts()
	if cachedPorts != nil {
		l.Info("Returning cached ports")
		*ports = cachedPorts
		return nil
	}
	//In case our cached ports have gone stale, we update them
	*ports = libs.FetchPorts()
	cacheAllPorts(*ports)
	return nil
}

// GetPort fetches the port with the specified port number
func GetPort(port *m.PortCBP, portNumber string) (err error) {
	cachedPort := getCachedPort(portNumber)
	if cachedPort != nil {
		l.Info("Returning cached port for the given [portNumber %v]", portNumber)
		port = cachedPort
		return nil
	}

	if err := portClient.Child(portNumber).Get(ctx, &port); err != nil {
		l.Error("Error fetching port #%v", err, portNumber)
		return err
	}
	if port == nil {
		return fmt.Errorf("Port #%s not found", portNumber)
	}
	cachePort(portNumber, port)
	return nil
}

// GetPortsByBorder returns a list of ports with the specified border name
func GetPortsByBorder(ports *[]m.PortCBP, border string) (err error) {
	cachedPorts := getCachedPortsByBorder(border)
	if cachedPorts != nil {
		l.Info("Returning cached ports for the given [border %v]", border)
		ports = cachedPorts
		return nil
	}

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
	cachePortsByBorder(border, ports)
	return nil
}

func getCachedPortsByBorder(border string) (ports *[]m.PortCBP) {
	res, found := cache.Get(portsByBorderKeyPrefix + border)
	if !found {
		l.Info("There are no ports cached for the given [border %v]", border)
		return nil
	}

	l.Info("Ports cached for the given [border %v]", border)
	return res.(*[]m.PortCBP)
}

func cachePortsByBorder(border string, ports *[]m.PortCBP) {
	key := portsByBorderKeyPrefix + border
	cache.Put(key, ports)
	l.Info("Successfully cached the ports by the [border %v]", border)
}

func getCachedPort(portNumber string) (port *m.PortCBP) {
	res, found := cache.Get(portCachedKeyPrefix + portNumber)
	if !found {
		l.Info("There is no port cached for the given [portNumber %v]", portNumber)
		return nil
	}

	l.Info("Port cached for the given [portNumber %v]", portNumber)
	return res.(*m.PortCBP)
}

func cachePort(portNumber string, port *m.PortCBP) {
	key := portCachedKeyPrefix + portNumber
	cache.Put(key, port)
	l.Info("Successfully cached the port by the [portNumber %v]", portNumber)
}

func getAllCachedPorts() (ports *[]m.PortCBP) {
	res, found := cache.Get(allPortsCachedKey)
	if !found {
		l.Info("There are no ports cached")
		return nil
	}

	l.Info("Ports cached")
	return res.(*[]m.PortCBP)
}

func cacheAllPorts(ports *[]m.PortCBP) {
	cache.Put(allPortsCachedKey, ports)
	l.Info("Successfully cached all the given ports")
}
