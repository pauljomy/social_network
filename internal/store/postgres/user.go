package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/models"
)

type userStore struct {
	db *pgxpool.Pool
}

func (s *userStore) Create(ctx context.Context, user *models.User) error {

	query := `INSERT INTO users (name, email, created_at)
				VALUES($1, $2, $3) RETURNING id, created_at`

	err := s.db.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
		user.CreatedAt,
	).Scan(
		&user.ID,
		&user.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (s *userStore) FindByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}
