package repository

import (
	"github.com/pintoter/basic-crud-books/internal/model"
	"context"
	"database/sql"
	"fmt"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) Create(ctx context.Context, user model.User) (int, error) {
	var (
		id    int
		query string = fmt.Sprintf("INSERT INTO %s (login, email, password) VALUES ($1, $2, $3) RETURNING id", collectionUsers)
	)

	err := u.db.QueryRowContext(ctx, query, user.Login, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UsersRepository) GetByID(ctx context.Context, id int) (model.User, error) {
	var (
		user  model.User
		query string = fmt.Sprintf("SELECT FROM %s WHERE id = $1", collectionUsers)
	)

	err := u.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Login, &user.Password, &user.RegisteredAt) // зачем мы тут сканируем ID, если нам человек и так отправляет айди
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UsersRepository) GetByLogin(ctx context.Context, login string) (model.User, error) {
	var (
		user  model.User
		query string = fmt.Sprintf("SELECT FROM %s WHERE login = $1", collectionUsers)
	)

	err := u.db.QueryRowContext(ctx, query, login).Scan(&user.ID, &user.Email, &user.Login, &user.Password, &user.RegisteredAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UsersRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var (
		user  model.User
		query string = fmt.Sprintf("SELECT FROM %s WHERE login = $1", collectionUsers)
	)

	err := u.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Login, &user.Password, &user.RegisteredAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UsersRepository) GetByCredentials(ctx context.Context, login, password string) (model.User, error) {
	var (
		user  model.User
		query string = fmt.Sprintf("SELECT FROM %s WHERE login = $1 AND password = $2", collectionUsers)
	)

	err := u.db.QueryRowContext(ctx, query, login, password).Scan(&user.ID, &user.Email, &user.Login, &user.Password, &user.RegisteredAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
