package repository

import "yadro/domain"

type Weather interface {
	GetWeather(Location, DateFrom, DateTo string) (*domain.WeatherResponse, error)
}
