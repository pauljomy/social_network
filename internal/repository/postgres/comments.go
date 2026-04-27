package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/models"
	"github.com/pauljomy/social/internal/repository"
)

type postgresCommentRepo struct {
	db *pgxpool.Pool
}

func NewCommentRepository(db *pgxpool.Pool) repository.CommentRepository {
	return &postgresCommentRepo{db: db}
}

func (r *postgresCommentRepo) Create(ctx context.Context, comment *models.Comment) error {
	return nil
}
