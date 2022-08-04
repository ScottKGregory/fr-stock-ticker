package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/scottkgregory/fr-stock-ticker/models"
)

var ok = &models.Probe{
	Status: "ok",
}

// Liveness should be used to confirm the application is operational. In this version
// it simply returns a static response but could be expanded depending on future requirements
func Liveness(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeData(log, w, ok)
	}
}

// Readiness should be used to confirm the application is ready to accept traffic. In this
// version it simply returns a static response but could be expanded depending on future requirements
func Readiness(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeData(log, w, ok)
	}
}
