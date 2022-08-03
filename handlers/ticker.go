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

func filterDays(av map[string]models.AvTimeSeries, days int) ([]models.FrTimeSeries, error) {
	ret := []models.FrTimeSeries{}
	for k, v := range av {
		date, err := time.Parse("2006-01-02", k)
		if err != nil {
			return nil, fmt.Errorf("error parsing date: %w", err)
		}

		frTimeSeries, err := models.FrTimeSeriesFromAvTimeSeries(&v, date)
		if err != nil {
			return nil, fmt.Errorf("error mapping av time series to fr time series")
		}

		ret = append(ret, *frTimeSeries)
	}

	sort.Slice(ret, func(i, j int) bool { return ret[i].Date.Before(ret[j].Date) })

	ret = ret[len(ret)-days:]

	return ret, nil
}

func calcAvgClose(ts []models.FrTimeSeries) (total float64) {
	if len(ts) == 0 {
		return 0
	}

	for _, t := range ts {
		total += t.Close
	}

	return total / float64(len(ts))
}
