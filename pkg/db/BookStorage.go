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

func ListBooks(db *pg.DB) ([]*models.Book, error) {
	var books []*models.Book

	err := db.Model(&books).Select()
	return books, err
}

func GetBook(db *pg.DB, id string) (*models.Book, error) {
	book := &models.Book{}

	err := db.Model(book).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return book, nil
}

func CreateBook(db *pg.DB, book models.Book) error {
	_, err := db.Model(book).Insert()
	if err != nil {
		return err
	}
	return nil
}
