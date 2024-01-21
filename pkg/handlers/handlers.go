package handlers

import (
	"encoding/json"
	"go-api/pkg/db"
	"net/http"

	"github.com/go-pg/pg"
)

type BookHandler struct {
	DB *pg.DB
}

// or, in main.go you could write:
// store := BookStore{}
// handler := BookHandler(storeage: store)
// allowing you to call, for example, `handler.list()`

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {

	if b.DB == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	books, err := db.List(b.DB)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

// func (b BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	book := b.storage.Get(id)
// 	if book == nil {
// 		http.Error(w, "Book not found", http.StatusNotFound)
// 	}
// 	err := json.NewEncoder(w).Encode(book)
// 	if err != nil {
// 		http.Error(w, "Internal error", http.StatusInternalServerError)
// 		return
// 	}
// }

// func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
// 	var book models.Book
// 	err := json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	b.storage.Create(book)
// 	err = json.NewEncoder(w).Encode(book)
// 	if err != nil {
// 		http.Error(w, "Internal error", http.StatusInternalServerError)
// 		return
// 	}
// }

// func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	var book models.Book
// 	err := json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	updatedBook := b.storage.Update(id, book)
// 	if updatedBook == nil {
// 		http.Error(w, "Book not found", http.StatusNotFound)
// 		return
// 	}
// 	err = json.NewEncoder(w).Encode(updatedBook)
// 	if err != nil {
// 		http.Error(w, "Internal error", http.StatusInternalServerError)
// 		return
// 	}
// }

// func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	book := b.storage.Delete(id)
// 	if book == nil {
// 		http.Error(w, "Book not found", http.StatusNotFound)
// 		return
// 	}

// }
