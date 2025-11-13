package report

import (
	"context"
	"fmt"
	"project_sem/internal/database"

	"github.com/google/uuid"
)

type Repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	repository := &Repository{db: db}

	return repository
}

func (r *Repository) Accept(parentCtx context.Context, UUID uuid.UUID) error {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	if _, err := r.db.Exec(ctx, "DELETE FROM reports"); err != nil {
		return fmt.Errorf("conn.Exec(DELETE): %w", err)
	}

	sql := `
INSERT INTO reports
	(uuid, id, name, category, price, create_date)
SELECT
	MIN(uuid), MIN(id), name, category, price, create_date
FROM
	prices
WHERE
	uuid=$1
GROUP BY 
	name, category, price, create_date`

	if _, err := r.db.Exec(ctx, sql, UUID.String()); err != nil {
		return fmt.Errorf("conn.Exec(INSERT ... SELECT): %w", err)
	}

	return nil
}
