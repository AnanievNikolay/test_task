package model

import (
	"encoding/json"
	"log"
)

//ExternalAPIResponse ...
type ExternalAPIResponse struct {
	Raw     map[string]map[string]*Currency `json:"RAW"`
	Display map[string]map[string]*Currency `json:"DISPLAY"`
}

//Marshal ..
func (ear *ExternalAPIResponse) Marshal() string {
	jObj, jErr := json.Marshal(ear)
	if jErr != nil {
		log.Fatal("[Error] Something went wrong while marshaling ExternalApiResponse. Error:", jErr.Error())
	}
	return string(jObj)
}
