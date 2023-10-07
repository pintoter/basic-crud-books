package model

import "errors"

var (
	ErrBookAlreadyExists = errors.New("Book already exists")
	ErrInvalidData       = errors.New("Your inputting data is invalid")
	ErrBooksNotFound     = errors.New("List of books is empty")
	ErrBookNotFound      = errors.New("Book with input ID doesn't exist")
	ErrUserLoginExist    = errors.New("User with input Login already exist")
	ErrUserEmailExist    = errors.New("User with input Email already exist")
)
