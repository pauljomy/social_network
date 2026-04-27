package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/repository"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Posts() repository.PostRepository {
	return &postgresPostRepo{db: r.db}
}

func (r *postgresRepository) Comments() repository.CommentRepository {
	return &postgresCommentRepo{db: r.db}
}

func (r *postgresRepository) Users() repository.UserRepository {
	return &postgresUserRepo{db: r.db}
}
