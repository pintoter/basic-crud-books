package repository

import (
	"github.com/pintoter/basic-crud-books/internal/model"
	"context"
	"database/sql"
)

type Books interface {
	Create(ctx context.Context, book model.Book) (int, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, id int) (model.Book, error)
	GetByAuthor(ctx context.Context, author string) ([]model.Book, error)
	GetByTitle(ctx context.Context, title string) (model.Book, error)
	UpdateRating(ctx context.Context, rating float64, author string, title string) error
	DeleteByID(ctx context.Context, id int) error
	DeleteByTitle(ctx context.Context, title string) error
}

type Users interface {
	Create(ctx context.Context, user model.User) (int, error)
	GetByID(ctx context.Context, id int) (model.User, error)
	GetByLogin(ctx context.Context, login string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetByCredentials(ctx context.Context, login, password string) (model.User, error)
	// GetByRefreshToken(ctx context.Context, Token string) (model.User, error)
	// SetSession(ctx context.Context, userId int, session model.Session) error
}

type Repository struct {
	Books
	Users
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Books: NewBooks(db),
		Users: NewUsers(db),
	}
}
