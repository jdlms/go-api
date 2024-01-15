package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	storage BookStorage
}

// or, in main.go you could write:
// store := BookStore{}
// handler := BookHandler(storeage: store)
// allowing you to call, for example, `handler.list()`

func (b BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(b.storage.List())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book := b.storage.Get(id)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.storage.Create(book)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBook := b.storage.Update(id, book)
	if updatedBook == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (b BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book := b.storage.Delete(id)
	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

}
