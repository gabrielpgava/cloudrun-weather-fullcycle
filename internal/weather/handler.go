package wheather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type weatherHandler struct {
	weatherService WeatherService
	cepResolver    CEPResolver
}

func NewWeatherHandler(weatherService WeatherService, cepResolver CEPResolver) *weatherHandler {
	return &weatherHandler{
		weatherService: weatherService,
		cepResolver:    cepResolver,
	}
}

func (h *weatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	if cep == "" {
		fmt.Println("CEP is missing")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "cep is required"})
		return
	}

	if len(cep) != 8 {
		fmt.Println("Invalid CEP length:", len(cep))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid zipcode"})
		return
	}

	cityName, err := h.cepResolver.GetCityByCep(cep)
	if err != nil {
		fmt.Println("Error resolving CEP:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "can not find zipcode"})
		return
	}

	weatherData, err := h.weatherService.GetWeather(cityName)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "error fetching weather data"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weatherData)
}
