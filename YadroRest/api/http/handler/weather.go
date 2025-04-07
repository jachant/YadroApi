package handler

import (
	"encoding/json"
	"net/http"
	"yadro/api/http/types"
	"yadro/usecases"

	"github.com/go-chi/chi/v5"
)

type Weather struct {
	service usecases.Weather
}

func NewWeatherHandler(service usecases.Weather) *Weather {
	return &Weather{service: service}
}

func (we *Weather) WithWeatherHandlers(r chi.Router) {
	r.Get("/weather", we.GetWeather)
}

func (we *Weather) GetWeather(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetWeatherRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //400
		return
	}
	resp, err := we.service.GetWeather(req.Location, req.DateFrom, req.DateTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //500
		return
	}
	result, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //500
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header()
	w.Write(result)

}
