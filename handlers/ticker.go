package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/rs/zerolog"
	"github.com/scottkgregory/fr-stock-ticker/models"
	"github.com/scottkgregory/fr-stock-ticker/services"
)

// Ticker is the primary handler for the API, it calls through to the supplied Alphavantage service to receive all available
// time series data before sorting and filtering the response to only the requested days
func Ticker(log zerolog.Logger, service services.AlphavantageOperations, days int, symbol string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := service.GetDailyTimeSeries(r.Context(), symbol)
		if err != nil {
			writeError(log, w, "error getting alphavantage data: %s", err)
			return
		}

		frTs, err := filterDays(resp.TimeSeries, days)
		if err != nil {
			writeError(log, w, "error filtering days: %s", err)
			return
		}

		writeData(log, w, &models.FrResponse{
			Symbol:       symbol,
			Days:         days,
			TimeSeries:   frTs,
			AverageClose: calcAvgClose(frTs),
		})
	}
}

// Filter days processes an Alphavantage response set and pulls only the last N days from it, sorting them before returning
func filterDays(av map[string]models.AvTimeSeries, days int) ([]models.FrTimeSeries, error) {
	ret := []models.FrTimeSeries{}
	for k, v := range av {
		date, err := time.Parse("2006-01-02", k)
		if err != nil {
			return nil, fmt.Errorf("error parsing date: %w", err)
		}

		frTimeSeries, err := models.FrTimeSeriesFromAvTimeSeries(v, date)
		if err != nil {
			return nil, fmt.Errorf("error mapping av time series to fr time series: %w", err)
		}

		ret = append(ret, *frTimeSeries)
	}

	sort.Slice(ret, func(i, j int) bool { return ret[i].Date.Before(ret[j].Date) })

	ret = ret[len(ret)-days:]

	return ret, nil
}

// calcAvgClose calculates the average of the close values for each time series in the slice supplied
func calcAvgClose(ts []models.FrTimeSeries) (total float64) {
	if len(ts) == 0 {
		return 0
	}

	for _, t := range ts {
		total += t.Close
	}

	return total / float64(len(ts))
}
