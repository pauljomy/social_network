package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/models"
)

type userStore struct {
	db *pgxpool.Pool
}

func (s *userStore) Create(ctx context.Context) error {
	return nil
}

func (s *userStore) FindByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}
