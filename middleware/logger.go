package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func Logger(log zerolog.Logger, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		l := log.With().Str("path", r.URL.Path).Logger()

		l.Info().Msg("request made")

		handler(w, r)

		l.Info().Dur("request-duration", time.Now().Sub(start)*time.Nanosecond).Msg("response returned")
	}
}
