package model

//Currency ...
type Currency struct {
	CHANGE24HOUR    interface{} `json:"CHANGE24HOUR,string"`
	CHANGEPCT24HOUR interface{} `json:"CHANGEPCT24HOUR,string"`
	OPEN24HOUR      interface{} `json:"OPEN24HOUR,string"`
	VOLUME24HOUR    interface{} `json:"VOLUME24HOUR,string"`
	VOLUME24HOURTO  interface{} `json:"VOLUME24HOURTO,string"`
	LOW24HOUR       interface{} `json:"LOW24HOUR,string"`
	HIGH24HOUR      interface{} `json:"HIGH24HOUR,string"`
	PRICE           interface{} `json:"PRICE,string"`
	SUPPLY          interface{} `json:"SUPPLY,string"`
	MKTCAP          interface{} `json:"MKTCAP,string"`
}
