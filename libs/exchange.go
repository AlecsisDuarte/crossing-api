package libs

import (
	"crossing-api/libs/cache"
	"crossing-api/libs/log"
	"crossing-api/models"
	"encoding/json"

	"net/http"
	"net/url"
)

const (
	baseURL                string = "https://api.exchangeratesapi.io/latest?"
	baseKey                string = "base"
	baseValue              string = "USD"
	symbolsKey             string = "symbols"
	mexAndCadSymbols       string = "MXN,CAD"
	exchangeDescription    string = "Current and historical foreign exchange rates published by the European Central Bank"
	exchangePrefixCacheKey string = "EXCHANGE_CACHED_"
)

// FetchExchangeRate makes a request to api.exchangeratesapi.io to fetch the specifide symbols exchanges rates using USD as base
func FetchExchangeRate(symbol string) *models.Exchange {
	cachedExchange := getCachedExchange(symbol)
	if cachedExchange != nil {
		return cachedExchange
	}

	url, err := url.Parse(baseURL)
	if err != nil {
		log.Error("Error parsing the exchange url %v", err, baseURL)
		return nil
	}

	query := url.Query()
	query.Set(baseKey, baseValue)
	query.Add(symbolsKey, symbol)
	url.RawQuery = query.Encode()
	res, err := http.Get(url.String())

	if err != nil {
		log.Error("Error fetching the exchanges", err)
		return nil
	}
	log.Info("Succesfully fetched the exchange rates")
	var exchange models.Exchange
	json.NewDecoder(res.Body).Decode(&exchange)
	exchange.Source = url.String()
	exchange.Description = exchangeDescription
	cacheExchange(symbol, &exchange)
	return &exchange
}

// FetchAllExchangeRates makes a request to api.exchangeratesapi.io to fetch both MXN and CAD exchanges rates using USD as base
func FetchAllExchangeRates() *models.Exchange {
	return FetchExchangeRate(mexAndCadSymbols)
}

func cacheExchange(symbol string, exchange *models.Exchange) {
	cache.Put(exchangePrefixCacheKey+symbol, exchange)
}

func getCachedExchange(symbol string) (exchange *models.Exchange) {
	res, found := cache.Get(exchangePrefixCacheKey + symbol)
	if !found {
		log.Info("There is no exchange cached")
		return nil
	}

	log.Info("Exchange cached")
	return res.(*models.Exchange)
}
