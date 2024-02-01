package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// format to use url
const weatherUrl = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s"

type WeatherService struct {
}

type WeatherResponse struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}

func newWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetWeatherOfCity(apiKey, city string) (*WeatherResponse, error) {
	resp, err := http.Get(fmt.Sprintf(weatherUrl, city, apiKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var d WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}
