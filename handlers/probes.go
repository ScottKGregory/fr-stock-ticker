package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/scottkgregory/fr-stock-ticker/models"
)

var ok = &models.Probe{
	Status: "ok",
}

func Liveness(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeData(log, w, ok)
	}
}

func Readiness(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeData(log, w, ok)
	}
}
