package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/models"
	"github.com/pauljomy/social/internal/repository"
)

type postgresPostRepo struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) repository.PostRepository {
	return &postgresPostRepo{db: db}
}

func (r *postgresPostRepo) Create(ctx context.Context, post *models.Post) error {
	return nil
}
