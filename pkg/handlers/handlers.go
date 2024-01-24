package handlers

import (
	"encoding/json"
	"go-api/pkg/db"
	"go-api/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg"
)

type BookHandler struct {
	DB *pg.DB
}

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {

	books, err := db.ListBooks(b.DB)
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

func (b BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := db.GetBook(b.DB, id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateBook(b.DB, newBook)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

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
