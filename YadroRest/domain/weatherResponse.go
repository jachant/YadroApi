package domain

type WeatherResponse struct {
    Days []Day `json:"days"`
}

// Данные за день
type Day struct {
    Date  string `json:"datetime"`       // Дата в формате "2023-10-25"
    Hours []Hour `json:"hours"`          // Почасовые данные
}

// Почасовые данные
type Hour struct {
    Time        string  `json:"datetime"`    // Время в формате "2023-10-25T14:00:00"
    Temperature float64 `json:"temp"`        // Температура
    Humidity    float64 `json:"humidity"`    // Влажность (пример дополнительного параметра)
}