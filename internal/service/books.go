package service

import (
	"books-app/internal/model"
	"books-app/internal/repository"
	"context"
	"database/sql"
	"errors"
)

type BooksService struct {
	booksRepo repository.Books
}

func NewBooks(booksRepo repository.Books) *BooksService {
	return &BooksService{
		booksRepo: booksRepo,
	}
}

func (b *BooksService) Create(ctx context.Context, book model.Book) (int, error) {
	isExists, err := b.isBookExists(ctx, book.Title)
	if err != nil {
		return 0, err
	}

	if isExists {
		return 0, model.ErrBookAlreadyExists
	}

	if (book.PublicationYear > 2023 || book.PublicationYear < 1) || (book.Rating > 5 || book.Rating < 1) {
		return 0, model.ErrInvalidData
	}

	return b.booksRepo.Create(ctx, book)
}

func (b *BooksService) GetAll(ctx context.Context) ([]model.Book, error) {
	books, err := b.booksRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, model.ErrBooksNotFound
	}

	return books, nil
}

func (b *BooksService) GetByID(ctx context.Context, id int) (model.Book, error) {
	book, err := b.booksRepo.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, model.ErrBookNotFound
		}

		return model.Book{}, err
	}

	return book, nil
}

func (b *BooksService) GetByAuthor(ctx context.Context, title string) ([]model.Book, error) {
	books, err := b.booksRepo.GetByAuthor(ctx, title)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrBookNotFound
		}

		return nil, err
	}

	if len(books) == 0 {
		return nil, model.ErrBooksNotFound
	}

	return books, nil
}

func (b *BooksService) UpdateBook(ctx context.Context, rating float64, author string, title string) error {
	_, err := b.isBookExists(ctx, title)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrBookNotFound
		}

		return err
	}

	return b.booksRepo.UpdateRating(ctx, rating, author, title)
}

func (b *BooksService) DeleteByID(ctx context.Context, id int) error {
	_, err := b.booksRepo.GetByID(ctx, id) // а если айди не существует как проверить?

	if err != nil {
		// добавил сегодня, добавить в работу с дб возврат несуществующей ячейки
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrBookNotFound
		}

		return err
	}

	return b.booksRepo.DeleteByID(ctx, id)
}

func (b *BooksService) DeleteByTitle(ctx context.Context, title string) error {
	isExists, err := b.isBookExists(ctx, title)

	if err != nil {
		return err
	}

	if !isExists {
		return model.ErrBookNotFound
	}

	return b.booksRepo.DeleteByTitle(ctx, title)
}

func (b *BooksService) isBookExists(ctx context.Context, title string) (bool, error) {
	_, err := b.booksRepo.GetByTitle(ctx, title)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
