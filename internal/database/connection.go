package database

import (
	"context"
	"fmt"
	"project_sem/internal/config"

	"github.com/jackc/pgx/v5"
)

var connection *pgx.Conn

func Connection() *pgx.Conn {
	if connection == nil {
		db, err := pgx.Connect(context.Background(), config.DataSourceName())
		if err != nil {
			panic(fmt.Errorf("pgx.Connect: %w", err))
		}

		connection = db
	}

	return connection
}
