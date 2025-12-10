package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postres, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Psql.User,
		cfg.Psql.Password,
		cfg.Psql.Host,
		cfg.Psql.Port,
		cfg.Psql.DBName,
	)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-1] Failed to connect to Postgres database " + cfg.Psql.Host)
		return nil, err
	}

	sqlDb, err := db.DB()
	
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-2] Failed to get database instance")
		return nil, err
	}

	sqlDb.SetMaxOpenConns(cfg.Psql.DBMaxOpen)
	sqlDb.SetMaxIdleConns(cfg.Psql.DBMaxIdle)

	return &Postres{DB: db}, nil
}