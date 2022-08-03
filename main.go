package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/scottkgregory/fr-stock-ticker/handlers"
	"github.com/scottkgregory/fr-stock-ticker/middleware"
	"github.com/scottkgregory/fr-stock-ticker/services"
)

func main() {
	log := zerolog.New(os.Stderr).With().Str("tag", "fr-stock-ticker").Logger()

	days, symbol, apiKey, err := getConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("invalid config supplied to application")
	}

	avService := services.NewAlphavantageService(log, 10, apiKey)

	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware.Logger(log, handlers.Ticker(log, avService, days, symbol)))
	mux.HandleFunc("/liveness", middleware.Logger(log, handlers.Liveness(log)))
	mux.HandleFunc("/readiness", middleware.Logger(log, handlers.Readiness(log)))

	log.Info().Msg("running webserver")
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("error running webserver")
	}
}

func getConfig() (int, string, string, error) {
	daysStr, daysOk := os.LookupEnv("FR_DAYS")
	if !daysOk {
		return 0, "", "", errors.New("FR_DAYS environment variable must be set")
	}

	days, err := strconv.Atoi(daysStr)
	if err != nil {
		return 0, "", "", fmt.Errorf("FR_DAYS environment variable must be a valid integer: %s", err)
	}

	symbol, symbolOk := os.LookupEnv("FR_SYMBOL")
	if !symbolOk {
		return 0, "", "", errors.New("FR_SYMBOL environment variable must be set")
	}

	apiKey, apiKeyOk := os.LookupEnv("FR_API_KEY")
	if !apiKeyOk {
		return 0, "", "", errors.New("FR_API_KEY environment variable must be set")
	}

	return days, symbol, apiKey, err
}
