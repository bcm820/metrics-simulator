package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"github.com/bcmendoza/metrics-simulator/metrics"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("service", "metrics-simulator").Logger()
	logger.Info().Msg("Starting metrics simulator")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	var servers []*http.Server
	viper.SetDefault("SERVICE_AMOUNT", 5)
	serviceAmount := viper.GetInt("SERVICE_AMOUNT")

	for i := 1; i <= serviceAmount; i++ {
		port := fmt.Sprintf("300%d", i)
		server := http.Server{
			Addr:    fmt.Sprintf("0.0.0.0:%s", port),
			Handler: handlers(port, logger, metrics.New()),
		}
		go func() {
			logger.Info().Msg(fmt.Sprintf("starting server on PORT %s", port))
			if err := server.ListenAndServe(); err != nil {
				logger.Error().AnErr("server.ListenAndServe", err).Msg(fmt.Sprintf("server on PORT %s", port))
			}
		}()
		servers = append(servers, &server)
	}

	s := <-sigChan
	logger.Info().Str("signal", s.String()).Msg("shutdown servers")
	for idx, server := range servers {
		port := fmt.Sprintf("300%d", idx+1)
		if err := server.Close(); err != nil {
			logger.Debug().AnErr("server.Close", err).Msg(fmt.Sprintf("shutdown server on PORT %s", port))
		} else {
			logger.Info().Msg(fmt.Sprintf("shutdown server on PORT %s", port))
		}
	}
}
