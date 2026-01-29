package main

import (
	"log"
	"net/http"
	"os"

	wheather "github.com/gabrielpgava/cloudrun-weather-fullcycle/internal/weather"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning: .env file not found, using environment variables")
	}

	httpClient := &http.Client{}
	repository := wheather.NewWeatherRepository(httpClient)
	weatherService := wheather.NewWeatherUseCase(repository)
	cepResolver := wheather.NewCEPResolver(repository)

	handler := wheather.NewWeatherHandler(weatherService, cepResolver)
	http.Handle("/weather", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("Starting server on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
