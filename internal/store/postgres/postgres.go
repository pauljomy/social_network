package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/store"
)

func NewPostStore(db *pgxpool.Pool) store.PostStorer {
	return &postStore{db: db}
}

func NewCommentStore(db *pgxpool.Pool) store.CommentStorer {
	return &commentStore{db: db}
}

func NewUserStore(db *pgxpool.Pool) store.UserStorer {
	return &userStore{db: db}
}
