package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/cfg"
	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/models"
)

func FreeWeatherRequest(city string) (*models.FreeWeatherPayload, error) {
	city = strings.ReplaceAll(strings.ToLower(city), " ", "+")
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", cfg.FreeWeather.ApiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
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
