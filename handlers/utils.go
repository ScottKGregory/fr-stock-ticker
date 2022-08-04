package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

var internalServerError []byte = []byte("internal server error")

// writeError should be used by handlers to write a generic error response. It
// will log specific details but not return them to the client
func writeError(log zerolog.Logger, w http.ResponseWriter, msg string, err error) {
	log.Error().Err(err).Msg(msg)
	w.WriteHeader(http.StatusInternalServerError)
	_, writeErr := w.Write(internalServerError)
	if writeErr != nil {
		log.Error().Err(writeErr).Msg("error writing failure response")
	}
}

// writeData JSON marshals the supplied data and writes it to the response. If any
// errors occur it will attempt to use writeError to return an error to the client
func writeData(log zerolog.Logger, w http.ResponseWriter, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		writeError(log, w, "error marshalling response", err)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		writeError(log, w, "error writing response", err)
	}
}
