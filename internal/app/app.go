package app

import (
	"bwanews/config"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()

	_, err := cfg.ConnectionPostgres()

	if err != nil {
		log.Fatal().Msgf("Error to connecting to database %v", err)
		return
	}
}