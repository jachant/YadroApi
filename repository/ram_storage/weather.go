package ram_storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"yadro/database"
	"yadro/domain"
)

type Weather struct {
	db *database.DataBase
}

func NewWeatherRepository(db *database.DataBase) *Weather {
	return &Weather{db: db}
}

func (we *Weather) GetWeather(Location, DateFrom, DateTo string) (*domain.WeatherResponse, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?key=%s", Location, DateFrom, DateTo, we.db.ApiKey)
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, errors.New(string(body))
	}
	Body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var weatherDays = domain.WeatherResponse{}
	if err := json.Unmarshal(Body, &weatherDays); err != nil {
		return nil, err
	}

	return &weatherDays, nil
}
