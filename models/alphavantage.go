package models

// AvResponse is the response object returned by the Alphavantage query API
type AvResponse struct {
	MetaData   AvMetaData              `json:"Meta Data"`
	TimeSeries map[string]AvTimeSeries `json:"Time Series (Daily)"`
}

// AvMetaData is the metadata format returned as part of the Alphavantage query API
type AvMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

// AvTimeSeries represents a single days time series data returned as part of the Alphavantage query API
type AvTimeSeries struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}
