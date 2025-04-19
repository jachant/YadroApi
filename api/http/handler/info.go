package handler

import (
	"encoding/json"
	"net/http"
	"yadro/usecases"

	"github.com/go-chi/chi/v5"
)

type Information struct {
	service usecases.Information
}

func NewInfoHndler(service usecases.Information) *Information {
	return &Information{service: service}
}

func (i *Information) WithInfoHandlers(r chi.Router) {
	r.Get("/", i.GetInfo)
}

func (i *Information) GetInfo(w http.ResponseWriter, r *http.Request) {
	info, err := i.service.GetInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // 500
	}
	resp, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //500
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
