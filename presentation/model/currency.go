package model

//Currency ...
type Currency struct {
	CHANGE24HOUR    interface{} `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR interface{} `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      interface{} `json:"OPEN24HOUR"`
	VOLUME24HOUR    interface{} `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  interface{} `json:"VOLUME24HOURTO"`
	LOW24HOUR       interface{} `json:"LOW24HOUR"`
	HIGH24HOUR      interface{} `json:"HIGH24HOUR"`
	PRICE           interface{} `json:"PRICE"`
	SUPPLY          interface{} `json:"SUPPLY"`
	MKTCAP          interface{} `json:"MKTCAP"`
}
