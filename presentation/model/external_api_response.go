package model

//ExternalApiResponse ...
type ExternalApiResponse struct {
	Raw     map[string]map[string]*Currency `json:"RAW"`
	Display map[string]map[string]*Currency `json:"DISPLAY"`
}
