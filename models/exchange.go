package models

// Exchange is a wrapper structure for data fetched from https://currencylayer.com/documentation
type Exchange struct {
	Quotes      	map[string]float64	 `json:"quotes"`
	Rates       	map[string]float64   `json:"rates,omitempty"`
	Source       	string             	`json:"source,omitempty"`
	Description		string             	`json:"description,omitempty"`
	URL      		string             	`json:"url,omitempty"`
	Symbol      	string             	`json:"symbol,omitempty"`
	Success   		bool				`json:"success,omitempty"`
	Terms 	    	string 				`json:"terms,omitempty"`
	Privacy   		string				`json:"privacy,omitempty"`
	Base   			string				`json:"base,omitempty"`
	Timestamp 		int    				`json:"timestamp,omitempty"`
}