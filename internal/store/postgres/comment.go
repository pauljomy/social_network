package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/models"
)

type commentStore struct {
	db *pgxpool.Pool
}

func (s *commentStore) Create(ctx context.Context, comment *models.Comment) error {
	return nil
}
