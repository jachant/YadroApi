package service

import (
	"sort"
	"yadro/domain"
	"yadro/repository"
	"yadro/usecases"
)

type Weather struct {
	repo repository.Weather
}

func NewWeatherService(repo repository.Weather) usecases.Weather {
	return &Weather{repo: repo}
}

func (we *Weather) GetWeather(Location, DateFrom, DateTo string) (*domain.Weather, error) {
	response, err := we.repo.GetWeather(Location, DateFrom, DateTo)
	if err != nil {
		return nil, err
	}
	Max := response.Days[0].Hours[0].Temperature
	Min := response.Days[0].Hours[0].Temperature
	Average := 0.0
	Median := 0.0
	temp := make([]float64, 0)
	for _, day := range response.Days {
		for _, hour := range day.Hours {
			if Max < hour.Temperature {
				Max = hour.Temperature
			}
			if Min > hour.Temperature {
				Min = hour.Temperature
			}
			Average += hour.Temperature
			temp = append(temp, hour.Temperature)
		}
	}
	sort.Float64s(temp)
	count := len(temp)
	Average /= float64(count)
	if mid := count / 2; mid%2 == 0 {
		Median = float64((temp[mid-1] + temp[mid]) / 2)
	} else {
		Median = float64(temp[mid-1])
	}
	Max = ((Max - 32) * 5) / 9
	Min = ((Min - 32) * 5) / 9
	Median = ((Median - 32) * 5) / 9
	Average = ((Average - 32) * 5) / 9

	result := domain.Weather{
		Data: struct {
			TemperatureC struct {
				Average float64 "json:\"average\""
				Median  float64 "json:\"median\""
				Min     float64 "json:\"min\""
				Max     float64 "json:\"max\""
			} "json:\"temperature_c\""
		}{struct {
			Average float64 "json:\"average\""
			Median  float64 "json:\"median\""
			Min     float64 "json:\"min\""
			Max     float64 "json:\"max\""
		}{Average: Average, Median: Median, Min: Min, Max: Max}},
		Service: "weather",
	}
	return &result, nil
}
