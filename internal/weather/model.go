package wheather

type WheatherData struct {
	CityName string  `json:"city_name"`
	Temp_C   float64 `json:"temp_c"`
	Temp_F   float64 `json:"temp_f"`
	Temp_k   float64 `json:"temp_k"`
}

func (w *WheatherData) ConvertToFahrenheit() float64 {
	return (w.Temp_C * 1.8) + 32
}

func (w *WheatherData) ConvertToKelvin() float64 {
	return w.Temp_C + 273.15
}
