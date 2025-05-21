package cep

import (
	"log"
	"net/http"

	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/handler"
	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/services"
	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/utils"
)

type ControllerGetCepTemperature struct {
	handler.TransactionControllerImpl
}

func (c *ControllerGetCepTemperature) Execute(payload interface{}) (result handler.ResponseController) {
	result = handler.NewJsonResponseController()
	cep := c.GetParam("cep")

	if !utils.ValidateCEP(cep) {
		log.Printf("invalid zipcode: %s", cep)
		result.SetResult(http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	cepPayload, _ := services.ViaCepRequest(cep)
	if cepPayload.Cep == "" {
		log.Printf("zipcode not found: %s", cep)
		result.SetResult(http.StatusNotFound, "zipcode not found")
		return
	}

	weatherPayload, err := services.FreeWeatherRequest(cepPayload.Localidade)
	if err != nil {
		log.Printf("error getting weather data: %s", err)
		result.SetResult(http.StatusInternalServerError, err)
		return
	}

	response := getCepTemperatureResponse{
		TemperatureCelsius:    weatherPayload.Current.TempC,
		TemperatureFahrenheit: weatherPayload.Current.TempF,
		TemperatureKelvin:     weatherPayload.Current.TempC + 273,
	}

	result.SetResult(http.StatusOK, response)
	return
}

type getCepTemperatureResponse struct {
	TemperatureCelsius    float64 `json:"temp_C"`
	TemperatureFahrenheit float64 `json:"temp_F"`
	TemperatureKelvin     float64 `json:"temp_K"`
}

func NewControllerGetCepTemperature() handler.Controller {
	return &ControllerGetCepTemperature{}
}
