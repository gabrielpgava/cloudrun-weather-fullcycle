package wheather

type weatherUseCase struct {
	repository Repository
}

func NewWeatherUseCase(repository Repository) *weatherUseCase {
	return &weatherUseCase{
		repository: repository,
	}
}

func (w *weatherUseCase) GetWeather(city string) (WheatherData, error) {
	data, err := w.repository.GetWeatherData(city)
	if err != nil {
		return WheatherData{}, err
	}
	data.Temp_F = data.ConvertToFahrenheit()
	data.Temp_k = data.ConvertToKelvin()
	return data, nil
}

type cepResolver struct {
	repository Repository
}

func NewCEPResolver(repository Repository) *cepResolver {
	return &cepResolver{
		repository: repository,
	}
}

func (c *cepResolver) GetCityByCep(cep string) (string, error) {
	return c.repository.GetCityByCep(cep)
}
