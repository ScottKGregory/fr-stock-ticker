package models

import (
	"fmt"
	"strconv"
	"time"
)

type FrResponse struct {
	Symbol       string         `json:"symbol"`
	Days         int            `json:"days"`
	TimeSeries   []FrTimeSeries `json:"timeSeries"`
	AverageClose float64        `json:"averageClose"`
}

type FrTimeSeries struct {
	Date   time.Time `json:"date"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume int       `json:"volume"`
}

func FrTimeSeriesFromAvTimeSeries(av *AvTimeSeries, date time.Time) (*FrTimeSeries, error) {
	open, err := strconv.ParseFloat(av.Open, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing float for open data: %w", err)
	}

	high, err := strconv.ParseFloat(av.High, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing float for high data: %w", err)
	}

	low, err := strconv.ParseFloat(av.Low, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing float for low data: %w", err)
	}

	close, err := strconv.ParseFloat(av.Close, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing float for close data: %w", err)
	}

	volume, err := strconv.Atoi(av.Volume)
	if err != nil {
		return nil, fmt.Errorf("error parsing int for volume data: %w", err)
	}

	return &FrTimeSeries{
		Date:   date,
		Open:   open,
		High:   high,
		Low:    low,
		Close:  close,
		Volume: volume,
	}, nil
}
