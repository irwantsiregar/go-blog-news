package app

import (
	"bwanews/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()

	_, err := cfg.ConnectionPostgres()

	if err != nil {
		log.Fatal().Msgf("Error to connecting to database %v", err)
		return
	}

	// CloudflareR2
	crfR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(crfR2)

}