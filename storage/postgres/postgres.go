package postgres

import (
	"context"
	"fmt"
	"log"
	"pulishla/config"
	"pulishla/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("dbname = %s password = %s port = %s host = %s user = %s sslmode=disable",
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
			cfg.PostgresPort,
			cfg.PostgresHost,
			cfg.PostgresUser,
		),
	)
	if err != nil {
		return nil, err
	}
	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Error while connecting to database !")
	}

	return &Store{db: pgxpool}, nil

}
