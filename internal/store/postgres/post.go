package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/models"
)

type postStore struct {
	db *pgxpool.Pool
}

func (s *postStore) Create(ctx context.Context, post *models.Post) error {

	query := `INSERT into posts (title,content, user_id, tags)
			    VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	err := s.db.QueryRow(ctx,
		query,
		post.Title,
		post.Content,
		post.UserID,
		post.Tags,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
