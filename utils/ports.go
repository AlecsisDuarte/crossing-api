package utils

import "crossing-api/models"

// ListToMap Changes the fethed list of CBP ports into a Map using the Port Number as the Key
func ListToMap(ports *[]models.PortCBP) map[string]models.PortCBP {
	portMaps := make(map[string]models.PortCBP)
	for i, port := range *ports {
		portMaps[port.PortNumber] = (*ports)[i]
	}
	return portMaps
}

// MapToList Changes the given map of CBP ports into a list of the values within
func MapToList(ports *map[string]models.PortCBP) []models.PortCBP {
	portsList := make([]models.PortCBP, len(*ports))
	index := 0
	for _, value := range *ports {
		portsList[index] = value
		index++
	}
	return portsList
}
