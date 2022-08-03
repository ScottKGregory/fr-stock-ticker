package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/scottkgregory/fr-stock-ticker/models"
)

type AlphavantageOperations interface {
	GetDailyTimeSeries(ct context.Context, symbol string) (*models.AvResponse, error)
}

type AlphavantageService struct {
	log        zerolog.Logger
	apiTimeout int
	apiKey     string
}

var _ AlphavantageOperations = &AlphavantageService{}

func NewAlphavantageService(log zerolog.Logger, apiTimeout int, apiKey string) AlphavantageOperations {
	return &AlphavantageService{log, apiTimeout, apiKey}
}

func (a *AlphavantageService) GetDailyTimeSeries(ct context.Context, symbol string) (*models.AvResponse, error) {
	ctx, cancel := context.WithTimeout(ct, time.Duration(a.apiTimeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s", a.apiKey, symbol),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error actioning request: %w", err)
	}

	if resp.Body == nil {
		return nil, fmt.Errorf("no body returned from http request")
	}

	parsedBody := models.AvResponse{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&parsedBody)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &parsedBody, nil
}
