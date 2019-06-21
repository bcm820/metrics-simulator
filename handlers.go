package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/bcmendoza/metrics-simulator/metrics"
)

func handlers(port string, logger zerolog.Logger, data *metrics.Metrics) http.Handler {
	logger = logger.With().Str("PORT", port).Logger()
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler(logger))
	r.HandleFunc("/metrics", mockMetricsHandler(logger, data))
	return r
}

// Hitting this route will update various metrics on the metrics object
func mainHandler(logger zerolog.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := checkMethod("/", r.Method, logger, w); ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte("Metrics simulator is live")); err != nil {
				logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
			} else {
				logger.Info().Msg("200 OK")
			}
		}
	}
}

// mockMetricsHandler responds with a fake metrics JSON object
func mockMetricsHandler(logger zerolog.Logger, data *metrics.Metrics) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := checkMethod("/metrics", r.Method, logger, w); ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			jsonResp, err := json.Marshal(data)
			if err != nil {
				logger.Error().AnErr("json.Marshal", err).Msg("500 Internal Server Error")
				Report(ProblemDetail{
					StatusCode: http.StatusInternalServerError,
					Detail:     "Could not marshall metrics into JSON",
				}, w)
			} else {
				w.Header().Set("Content-Type", "application/json")
				if _, err = w.Write(jsonResp); err != nil {
					logger.Error().AnErr("w.Write", err).Msg("500 Internal Server Error")
				} else {
					logger.Info().Msg("200 OK")
				}
			}
		}
	}
}

func checkMethod(route string, method string, logger zerolog.Logger, w http.ResponseWriter) (zerolog.Logger, bool) {
	logger = logger.With().Str("request-type", fmt.Sprintf("%s:'%s'", method, route)).Logger()
	if method != "GET" {
		logger.Warn().Msg("405 Method Not Allowed")
		Report(ProblemDetail{StatusCode: http.StatusMethodNotAllowed, Detail: method}, w)
		return logger, false
	}
	return logger, true
}
