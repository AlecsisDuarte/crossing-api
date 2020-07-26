package libs

import (
	"encoding/json"
	"net/http"
	"strings"

	"crossing-api/libs/log"
	"crossing-api/models"
)

const (
	cbpURL         string = "https://bwt.cbp.gov/api/bwtnew"
	mexicanBorder  string = "Mexican Border"
	canadianBorder string = "Canadian Border"
)

// FetchPorts fetches the latest status of the CBP ports
func FetchPorts() *[]models.PortCBP {
	log.Info("Calling CBP to get latest status of the ports")
	var ports = make([]models.PortCBP, 0)
	res, err := http.Get(cbpURL)
	if err != nil {
		log.Fatal("Error while fetching CBP info", err)
		return &ports
	}
	json.NewDecoder(res.Body).Decode(&ports)
	log.Info("Successfully fetch %d ports\n", len(ports))
	return &ports
}

// countryToBorderMap returns the map of the country to it's borders
func countryToBorderMap() map[string]string {
	return map[string]string{
		"mexico": mexicanBorder,
		"canada": canadianBorder,
		"mx":     mexicanBorder,
		"ca":     canadianBorder,
		"mex":    mexicanBorder,
		"can":    canadianBorder,
	}
}

// TranslateCountryToCBPBorder maps the contry to it's border name within the US or an empty string
// with a found value equal to false
func TranslateCountryToCBPBorder(country string) (border string, found bool) {
	lowerCaseCountry := strings.ToLower(country)
	countryToBorder := countryToBorderMap()
	border, found = countryToBorder[lowerCaseCountry]
	if !found {
		log.Info("The country %v is not a valid US border", country)
	}
	return border, found
}
