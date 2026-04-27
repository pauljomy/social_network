package repository

type Repository interface {
	Posts() PostRepository
	Comments() CommentRepository
	Users() UserRepository
}
