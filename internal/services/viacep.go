package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/models"
)

func ViaCepRequest(cep string) (*models.ViaCepPayload, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
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

	var result models.ViaCepPayload
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
