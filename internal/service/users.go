package service

import (
	"github.com/pintoter/basic-crud-books/internal/model"
	"github.com/pintoter/basic-crud-books/internal/repository"
	"github.com/pintoter/basic-crud-books/pkg/hash"
	"context"
	"database/sql"
	"errors"
)

type UsersService struct {
	usersRepo repository.Users
	hasher    hash.PasswordHasher
}

func NewUsers(users repository.Users, hasher hash.PasswordHasher) *UsersService {
	return &UsersService{
		usersRepo: users,
		hasher:    hasher,
	}
}

func (u *UsersService) SingUp(ctx context.Context, login, email, password string) (int, error) {
	if u.isLoginExists(ctx, login) {
		return 0, model.ErrUserLoginExist
	}

	if u.isEmailExists(ctx, email) {
		return 0, model.ErrUserEmailExist
	}

	id, err := u.usersRepo.Create(ctx, model.User{Login: login, Email: email, Password: password})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UsersService) SingIn(ctx context.Context, login, password string) (Tokens, error) {
	if !u.isLoginExists(ctx, login) {
		return Tokens{}, model.ErrUserLoginExist
	}

	// hashedPassword, err :=

	// id, err := u.usersRepo.GetByCredentials(ctx, login, password)
	// if err != nil {
	// 	return Tokens{}, err
	// }

	return Tokens{}, nil // исправить Tokens{}
}

func (u *UsersService) isLoginExists(ctx context.Context, login string) bool {
	_, err := u.usersRepo.GetByLogin(ctx, login)
	return !errors.Is(err, sql.ErrNoRows)
}

func (u *UsersService) isEmailExists(ctx context.Context, email string) bool {
	_, err := u.usersRepo.GetByEmail(ctx, email)
	return !errors.Is(err, sql.ErrNoRows)
}
