package service

import (
	"github.com/pintoter/basic-crud-books/internal/model"
	"github.com/pintoter/basic-crud-books/internal/repository"
	"github.com/pintoter/basic-crud-books/pkg/hash"
	"context"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Books interface {
	Create(ctx context.Context, book model.Book) (int, error)
	GetAll(ctx context.Context) ([]model.Book, error)
	GetByID(ctx context.Context, id int) (model.Book, error)
	GetByAuthor(ctx context.Context, title string) ([]model.Book, error)
	UpdateBook(ctx context.Context, rating float64, author string, title string) error
	DeleteByID(ctx context.Context, id int) error
	DeleteByTitle(ctx context.Context, title string) error
}

type Users interface {
	SingUp(ctx context.Context, login, email, password string) (int, error)
	SingIn(ctx context.Context, login, password string) (Tokens, error)
	// RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
}

type Service struct {
	Books
	Users
}

type Deps struct {
	Repos  *repository.Repository
	Hasher hash.PasswordHasher
}

func New(deps Deps) *Service {
	return &Service{
		Books: NewBooks(deps.Repos.Books),
		Users: NewUsers(deps.Repos.Users, deps.Hasher),
	}
}
