package libs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"crossing-api/models"
)

const cbpURL string = "https:// bwt.cbp.gov/api/bwtnew"

// FetchPorts fetches the latest status of the CBP ports
func FetchPorts() *[]models.PortCBP {
	log.Println("Calling CBP to get latest status of the ports")
	var ports = make([]models.PortCBP, 0)
	res, err := http.Get(cbpURL)
	if err != nil {
		log.Fatal(err)
		return &ports
	}
	json.NewDecoder(res.Body).Decode(&ports)
	fmt.Printf("Successfully fetch %d ports\n", len(ports))
	return &ports
}
