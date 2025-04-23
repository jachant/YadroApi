package usecases

import "yadro/domain"

type Weather interface {
	GetWeather(Location, DateFrom, DateTo string) (*domain.Weather, error)
}
