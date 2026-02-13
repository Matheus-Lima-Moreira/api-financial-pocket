package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/config"
	"github.com/rs/zerolog"
)

func main() {
  cfg := config.Load()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	logger.Info().
		Str("service", "financial-pocket").
		Str("port", cfg.Port).
		Msg("starting server")

	mux := http.NewServeMux()
	err := http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
