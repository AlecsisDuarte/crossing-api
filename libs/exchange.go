package libs

import (
	"crossing-api/libs/cache"
	"crossing-api/libs/log"
	"crossing-api/models"
	"crossing-api/utils"
	"encoding/json"

	"net/http"
	"net/url"
)

const (
	baseURL              		string = "http://api.currencylayer.com/live?"
	accessKeyKey				string = "access_key"
	sourceKey              		string = "source"
	baseValue              		string = "USD"
	currenciesKey          		string = "currencies"
	mexAndCadSymbols       		string = "MXN,CAD"
	exchangeDescription			string = "The currencylayer API is a product built and maintained by apilayer."
	exchangePrefixCacheKey	     string = "EXCHANGE_CACHED_"
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
	query.Set(sourceKey, baseValue)
	query.Add(currenciesKey, symbol)
	query.Add(accessKeyKey, utils.GetCurrencyLayerAccessKey())
	url.RawQuery = query.Encode()
	res, err := http.Get(url.String())

	if err != nil {
		log.Error("Error fetching the exchanges", err)
		return nil
	}
	log.Info("Succesfully fetched the exchange rates")
	var exchange models.Exchange
	json.NewDecoder(res.Body).Decode(&exchange)
	exchange.URL = url.String()
	exchange.Description = exchangeDescription
	//Legacy fields
	exchange.Rates = quotesToRates(&exchange)
	exchange.Base = exchange.Source
	cacheExchange(symbol, &exchange)
	return &exchange
}

// FetchAllExchangeRates makes a request to api.exchangeratesapi.io to fetch both MXN and CAD exchanges rates using USD as base
func FetchAllExchangeRates() *models.Exchange {
	return FetchExchangeRate(mexAndCadSymbols)
}

func cacheExchange(symbol string, exchange *models.Exchange) {
	cacheDuration := utils.GetExchangeCacheExpirationKey();
	cache.PutWithDuration(exchangePrefixCacheKey+symbol, exchange, cacheDuration)
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

func quotesToRates(exchange *models.Exchange) map[string]float64 {
	var rates = map[string]float64{}
	
	for key, value := range exchange.Quotes {
		rates[key[3:]] = value
	}
	return rates
}
