package libs

import (
	"crossing-api/models"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

const (
	baseURL      string = "https://api.exchangeratesapi.io/latest?"
	baseKey      string = "base"
	baseValue    string = "USD"
	symbolsKey   string = "symbols"
	symbolsValue string = "MXN,CAD"
)

// FetchExchangeRate makes a request to api.exchangeratesapi.io to fetch MXN and CAD exchanges rates using USD as base
func FetchExchangeRate() *models.Exchange {
	url, err := url.Parse(baseURL)
	if err != nil {
		log.Println("Error parsing the exchange url: ", err)
		return nil
	}

	query := url.Query()
	query.Set(baseKey, baseValue)
	query.Add(symbolsKey, symbolsValue)
	url.RawQuery = query.Encode()
	res, err := http.Get(url.String())

	if err != nil {
		log.Println("Error fetching the exchanges: ", err)
		return nil
	}
	log.Println("Succesfully fetched the exchange rates")
	var exchange models.Exchange
	json.NewDecoder(res.Body).Decode(&exchange)
	return &exchange
}
