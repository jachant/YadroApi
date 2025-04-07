package domain

type Weather struct {
	Data struct {
		TemperatureC struct {
			Average float64 `json:"average"`
			Median  float64 `json:"median"`
			Min     float64 `json:"min"`
			Max     float64 `json:"max"`
		} `json:"temperature_c"`
	} `json:"data"`
	Service string `json:"service"`
}
