package libs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"crossing-api/models"
)

const CBP_URL = "https://bwt.cbp.gov/api/bwtnew"

func FetchPorts() *[]models.PortCBP {
	log.Println("Calling CBP to get latest status of the ports")
	var ports = make([]models.PortCBP, 0)
	res, err := http.Get(CBP_URL)
	if err != nil {
		log.Fatal(err)
		return &ports
	}
	json.NewDecoder(res.Body).Decode(&ports)
	fmt.Printf("Successfully fetch %d ports\n", len(ports))
	return &ports
}
