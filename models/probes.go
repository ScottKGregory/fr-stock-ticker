package models

// Probes is a response model containing status data, used for liveness and readiness probes
type Probe struct {
	Status string `json:"status"`
}
