package models

type BookStorage interface {
	ListBooks() []*Book
	GetBooks(string) *Book
	UpdateBooks(string, Book) *Book
	CreateBooks(Book)
	DeleteBooks(string) *Book
}

type BookStore struct{}

type Book struct {
	ID               int    `pg:",pk" json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	PublishedDate    string `json:"published_date"`
	OriginalLanguage string `json:"original_language"`
}

var books = []*Book{
	{
		ID:               1,
		Title:            "7 habits of Highly Effective People",
		Author:           "Stephen Covey",
		PublishedDate:    "15/08/1989",
		OriginalLanguage: "English",
	},
}

// memory store + methods

func (b BookStore) List() []*Book {
	return books
}

func (b BookStore) Get(id int) *Book {
	for _, book := range books {
		if book.ID == id {
			return book
		}
	}
	return nil
}

func (b BookStore) Create(book Book) {
	books = append(books, &book)
}

func (b BookStore) Update(id int, bookUpdate Book) *Book {
	for i, book := range books {
		if book.ID == id {
			books[i] = &bookUpdate
			return book
		}
	}
	return nil
}

func (b BookStore) Delete(id int) *Book {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], (books)[i+1:]...)
			return &Book{}
		}
	}
	return nil
}
