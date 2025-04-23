package types

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

type GetWeatherRequest struct {
	Location string `json:"city" validate:"required"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

func (r *GetWeatherRequest) Validate() error {
	return Validator.Struct(r)
}

func CreateGetWeatherRequest(r *http.Request) (*GetWeatherRequest, error) {
	req := GetWeatherRequest{
		Location: r.URL.Query().Get("city"),
		DateFrom: r.URL.Query().Get("date_from"),
		DateTo:   r.URL.Query().Get("date_to"),
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	now := time.Now()
	if req.DateTo == "" {
		today := now.Truncate(24 * time.Hour)
		req.DateTo = today.Format("2006-01-02")
	}
	if req.DateFrom == "" {
		yesterday := now.AddDate(0, 0, -1).Truncate(24 * time.Hour)
		req.DateFrom = yesterday.Format("2006-01-02")
	}
	fmt.Println(req)
	return &req, nil
}
