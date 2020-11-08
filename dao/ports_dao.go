package dao

import (
	"crossing-api/libs"
	"crossing-api/libs/cache"
	l "crossing-api/libs/log"
	m "crossing-api/models"
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
