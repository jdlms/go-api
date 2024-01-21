package db

import (
	"go-api/pkg/models"

	"github.com/go-pg/pg"
)

// type BookStorage interface {
// 	List() []*Book
// 	Get(string) *Book
// 	Update(string, Book) *Book
// 	Create(Book)
// 	Delete(string) *Book
// }

type PostgresBookStore struct {
	db *pg.DB
}

func NewPostgresBookStore(db *pg.DB) *PostgresBookStore {
	return &PostgresBookStore{db: db}
}

func List(db *pg.DB) ([]*models.Book, error) {
	var books []*models.Book

	err := db.Model(&books).Select()
	return books, err
}

func Get(db *pg.DB, id string) (*models.Book, error) {
	book := &models.Book{}

	err := db.Model(book).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return book, nil
}
