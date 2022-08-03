package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

var internalServerError []byte = []byte("internal server error")

func writeError(log zerolog.Logger, w http.ResponseWriter, msg string, err error) {
	log.Error().Err(err).Msg(msg)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(internalServerError)
}

func writeData(log zerolog.Logger, w http.ResponseWriter, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		writeError(log, w, "error marshalling response", err)
		return
	}

	w.Write(body)
}
