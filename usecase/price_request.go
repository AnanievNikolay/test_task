package usecase

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/AnanievNikolay/test_task/domain"
	"github.com/AnanievNikolay/test_task/presentation/model"
)

//NewPriceRequest ..
func NewPriceRequest(_client domain.IClient) *PriceRequest {
	return &PriceRequest{
		client: _client,
	}
}

//PriceRequest ...
type PriceRequest struct {
	client domain.IClient
}

//Response ...
func (usecase *PriceRequest) Response() *model.ExternalAPIResponse {
	respModel := usecase.unmarshalJSON(usecase.client.Response())
	if reflect.DeepEqual(respModel, &model.ExternalAPIResponse{}) {
		log.Println("[PriceUseCase | Response] Empty model from usecase.")
		return nil
	}
	return respModel
}

func (usecase *PriceRequest) unmarshalJSON(_json string) *model.ExternalAPIResponse {
	obj := &model.ExternalAPIResponse{}
	err := json.Unmarshal([]byte(_json), obj)
	if err != nil {
		log.Println("[PriceUseCase | unmarshalJSON] Error whhile unmarshal. Error: ", err.Error())
	}
	return obj
}
