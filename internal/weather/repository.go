package wheather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

type weatherRepository struct {
	httpClient *http.Client
	weatherURL string
	cepURL     string
	apiKey     string
}

func NewWeatherRepository(httpClient *http.Client) *weatherRepository {
	return &weatherRepository{
		httpClient: httpClient,
		weatherURL: "http://api.weatherapi.com/v1/current.json",
		cepURL:     "https://viacep.com.br/ws",
		apiKey:     os.Getenv("WHEATHERKEY"),
	}
}

func (w *weatherRepository) GetWeatherData(city string) (WheatherData, error) {
	if w.apiKey == "" {
		return WheatherData{}, fmt.Errorf("WHEATHERKEY environment variable not set")
	}

	url := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", w.weatherURL, w.apiKey, city)
	resp, err := w.httpClient.Get(url)
	if err != nil {
		return WheatherData{}, err
	}
	defer resp.Body.Close()

	var body map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return WheatherData{}, err
	}

	weatherInfo := body["current"].(map[string]interface{})
	return WheatherData{
		CityName: city,
		Temp_C:   weatherInfo["temp_c"].(float64),
	}, nil
}

func (w *weatherRepository) GetCityByCep(cep string) (string, error) {
	url := fmt.Sprintf("%s/%s/json/", w.cepURL, cep)
	resp, err := w.httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var body map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return "", err
	}
	cityRaw := body["localidade"].(string)
	cityFormatted := formatCityName(cityRaw)
	return cityFormatted, nil
}

func formatCityName(s string) string {
	t := norm.NFD.String(strings.ToLower(strings.TrimSpace(s)))
	t = strings.Map(func(r rune) rune {
		switch {
		case unicode.Is(unicode.Mn, r):
			return -1
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			return r
		case unicode.IsSpace(r) || r == '-' || r == '_':
			return ' '
		default:
			return -1
		}
	}, t)

	return strings.Join(strings.Fields(t), "-")
}
