package books

import (
	"fmt"
	"log"

	"go-api/internal/models"

	"github.com/go-pg/pg"
)

type PostgresBookStore struct {
	db *pg.DB
}

func NewPostgresBookStore(db *pg.DB) *PostgresBookStore {
	return &PostgresBookStore{db: db}
}

// functions for in-memory store

func ListBooks(db *pg.DB) ([]*models.Book, error) {
	var books []*models.Book

	if db == nil {
		// Log the error and return
		log.Println("Failed to connect to the database")
		return nil, fmt.Errorf("failed to connect to the database")
	}

	err := db.Model(&books).Select()
	return books, err
}

func GetBook(db *pg.DB, id int) (*models.Book, error) {
	book := &models.Book{}

	err := db.Model(book).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return book, nil
}

func CreateBook(db *pg.DB, book *models.Book) error {
	newBook, err := db.Model(book).Insert()
	if err != nil {
		log.Printf("Error inserting book into database: %v", err)
		return (err)
	}
	fmt.Print(newBook)
	return nil
}

func UpdateBook(db *pg.DB, id int, book models.Book) error {
	existingBook := &models.Book{ID: id}
	err := db.Model(existingBook).WherePK().Select()
	if err != nil {
		return err
	}

	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.PublishedDate = book.PublishedDate
	existingBook.OriginalLanguage = book.OriginalLanguage

	_, err = db.Model(existingBook).WherePK().Update()
	return err
}

func DeleteBook(db *pg.DB, id int) error {
	book := &models.Book{ID: id}
	err := db.Delete(book)
	if err != nil {
		return err
	}
	return nil
}
