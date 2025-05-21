package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/cfg"
	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/models"
)

func FreeWeatherRequest(city string) (*models.FreeWeatherPayload, error) {
	city = url.QueryEscape(strings.ToLower(city))
	log.Printf("city: %s", city)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", cfg.FreeWeather.ApiKey, city)
	log.Printf("requesting FreeWeather API: %s", url)

	resp, err := http.Get(url)
	log.Printf("response FreeWeather API: %s", resp.Request.Body)
	if err != nil {
		log.Printf("error making request to FreeWeather API: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("error in request to FreeWeather API: %s", resp.Status)
		return nil, fmt.Errorf("erro na requisição: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result models.FreeWeatherPayload
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
