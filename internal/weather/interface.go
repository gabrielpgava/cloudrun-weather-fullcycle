package wheather

// WeatherService define operações de clima
type WeatherService interface {
	GetWeather(city string) (WheatherData, error)
}

// CEPResolver resolve CEP em cidade
type CEPResolver interface {
	GetCityByCep(cep string) (string, error)
}

// Repository abstrai a fonte de dados
type Repository interface {
	GetWeatherData(city string) (WheatherData, error)
	GetCityByCep(cep string) (string, error)
}
