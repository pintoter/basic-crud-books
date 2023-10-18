package repository

import (
	"github.com/pintoter/basic-crud-books/internal/model"
	"context"
	"database/sql"

	"fmt"
)

type BooksRepository struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *BooksRepository {
	return &BooksRepository{db: db}
}

func (b *BooksRepository) Create(ctx context.Context, book model.Book) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (author, book_name, rating, publication_date) VALUES ($1, $2, $3, $4) RETURNING id", collectionBooks)
	err := b.db.QueryRowContext(ctx, query, book.Author, book.Title, book.Rating, book.PublicationYear).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (b *BooksRepository) GetAll(ctx context.Context) ([]model.Book, error) {
	var bks []model.Book

	query := fmt.Sprintf("SELECT * FROM %s", collectionBooks)
	rows, err := b.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bk model.Book
		if err := rows.Scan(&bk.ID, &bk.Author, &bk.Title, &bk.Rating, &bk.PublicationYear); err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

func (b *BooksRepository) GetByID(ctx context.Context, id int) (model.Book, error) {
	bk := model.Book{}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", collectionBooks)
	err := b.db.QueryRowContext(ctx, query, id).Scan(&bk.ID, &bk.Author, &bk.Title, &bk.Rating, &bk.PublicationYear)

	if err != nil {
		return bk, err
	}

	return bk, nil
}

func (b *BooksRepository) GetByAuthor(ctx context.Context, author string) ([]model.Book, error) {
	var bks []model.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE author = $1", collectionBooks)
	rows, err := b.db.QueryContext(ctx, query, author)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bk model.Book
		if err := rows.Scan(&bk.ID, &bk.Author, &bk.Title, &bk.Rating, &bk.PublicationYear); err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

func (b *BooksRepository) GetByTitle(ctx context.Context, title string) (model.Book, error) {
	bk := model.Book{}

	query := fmt.Sprintf("SELECT * FROM %s WHERE book_name = $1", collectionBooks)
	err := b.db.QueryRowContext(ctx, query, title).Scan(&bk.ID, &bk.Author, &bk.Title, &bk.Rating, &bk.PublicationYear)
	if err != nil {
		return model.Book{}, err
	}

	return bk, nil
}

func (b *BooksRepository) UpdateRating(ctx context.Context, rating float64, author string, title string) error {
	query := fmt.Sprintf("UPDATE %s SET rating = $1 WHERE book_name = $2 AND author = $3", collectionBooks)

	_, err := b.db.ExecContext(ctx, query, rating, title, author)
	if err != nil {
		return err
	}

	return nil
}

func (b *BooksRepository) DeleteByID(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", collectionBooks)

	_, err := b.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (b *BooksRepository) DeleteByTitle(ctx context.Context, title string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name = $1", collectionBooks)

	_, err := b.db.ExecContext(ctx, query, title)
	if err != nil {
		return err
	}

	return nil
}
