# CEP Weather API

API em Go que retorna temperaturas (Celsius, Fahrenheit e Kelvin) para um CEP brasileiro válido.

## Funcionalidades

- Consulta temperatura via CEP (8 dígitos).
- Conversão de temperatura:
  - Fahrenheit: `C * 1.8 + 32`
  - Kelvin: `C + 273`

## Integrações Externas

- [ViaCEP](https://viacep.com.br/)
- [WeatherAPI](https://www.weatherapi.com/)

## Endpoints

- `GET /appstatus` – Status da aplicação
- `GET /cep/:cep` – Retorna temperatura
- `GET /swagger/index.html` – Documentação Swagger

## Respostas

- `200 OK`:
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

- `422` – CEP inválido (`invalid zipcode`)
- `404` – CEP não encontrado (`can not find zipcode`)

## Deploy (Google Cloud Run)

👉 https://goexpert-challenge-1-103094552570.us-central1.run.app/cep/:cep

## Execução local

**Pré-requisitos**: Docker, Docker Compose, chave da WeatherAPI.

**Passos**:
```bash
git clone https://github.com/HenriqueOtsuka/goexpert-challenge-1
cd seu-repo
# Adicione a API Key no docker-compose.yml
docker-compose up
```

Acesse:
- App: http://localhost:8080  
- Swagger: http://localhost:8080/swagger/index.html
