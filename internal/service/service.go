package service

type Weather interface {
	GetWeatherOfCity(apiKey, city string) (*WeatherResponse, error)
}

type Service struct {
	Weather
}

func NewService() *Service {
	return &Service{
		Weather: newWeatherService(),
	}
}
