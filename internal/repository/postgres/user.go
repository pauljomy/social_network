package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/models"
	"github.com/pauljomy/social/internal/repository"
)

type postgresUserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &postgresUserRepo{db: db}
}

func (r *postgresUserRepo) Create(ctx context.Context) error {
	return nil

}

func (r *postgresUserRepo) FindByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}
