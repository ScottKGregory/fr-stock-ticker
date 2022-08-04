package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

// Logger is a simple middleware used to log each request to the API along with it's request duration
func Logger(log zerolog.Logger, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		l := log.With().Str("path", r.URL.Path).Logger()

		l.Info().Msg("request made")

		handler(w, r)

		l.Info().Dur("request-duration", time.Since(start)).Msg("response returned")
	}
}
