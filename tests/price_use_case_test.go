package tests

import (
	"reflect"
	"testing"

	"github.com/AnanievNikolay/test_task/presentation/model"
	"github.com/AnanievNikolay/test_task/tests/client"
	"github.com/AnanievNikolay/test_task/usecase"
)

func TestPriceUseCase(t *testing.T) {
	testClient := &client.TestClient{}
	priceUseCase := usecase.NewPriceRequest(testClient)
	response := priceUseCase.Response()
	usd := response.Raw["BTC"]["USD"]
	shouldUsd := &model.Currency{
		CHANGE24HOUR:    float64(2603.709999999999),
		CHANGEPCT24HOUR: float64(5.826223162040298),
		OPEN24HOUR:      float64(44689.5),
		VOLUME24HOUR:    float64(111443.53408456499),
		VOLUME24HOURTO:  float64(5184795135.4528675),
		LOW24HOUR:       float64(44292.81),
		HIGH24HOUR:      float64(48219.85),
		PRICE:           float64(47293.21),
		SUPPLY:          float64(18624393),
		MKTCAP:          float64(880807329271.53),
	}
	if !reflect.DeepEqual(shouldUsd, usd) {
		t.Error("Should USD is not equal to USD")
	}
}
