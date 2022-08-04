package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog"
	"github.com/scottkgregory/fr-stock-ticker/models"
)

// ErrNoBody will be returned when the returned body is nil
var ErrNoBody = errors.New("no body returned from http request")

// AlphavantageOperations is the interface to which an Alphavantage service must adhere
type AlphavantageOperations interface {
	GetDailyTimeSeries(ct context.Context, symbol string) (*models.AvResponse, error)
}

// AlphavantageService is the primary implementation of an AlphavantageOperations interface
type AlphavantageService struct {
	log        zerolog.Logger
	apiTimeout int
	apiKey     string
}

var _ AlphavantageOperations = &AlphavantageService{}

// NewAlphavantageService assembles a new Alphavantage servise using the supplied config values
func NewAlphavantageService(log zerolog.Logger, apiTimeout int, apiKey string) *AlphavantageService {
	return &AlphavantageService{log, apiTimeout, apiKey}
}

// GetDailyTimeSeries calls through to the Alphavantage query API using the TIME_SERIES_DAILY function and the supplied symbol
func (a *AlphavantageService) GetDailyTimeSeries(ct context.Context, symbol string) (ret *models.AvResponse, err error) {
	ctx, cancel := context.WithTimeout(ct, time.Duration(a.apiTimeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(
			"https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s",
			url.QueryEscape(a.apiKey),
			url.QueryEscape(symbol),
		),
		http.NoBody,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error actioning request: %w", err)
	}

	if resp.Body == nil {
		return nil, ErrNoBody
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			a.log.Error().Err(closeErr).Msg("error closing response body")
			if err == nil {
				err = fmt.Errorf("error closing response body: %w", closeErr)
			}
		}
	}()

	parsedBody := models.AvResponse{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&parsedBody)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &parsedBody, nil
}
