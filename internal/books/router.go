package books

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg"
)

func Routes(pgdb *pg.DB) chi.Router {
	r := chi.NewRouter()

	bookHandler := BooksHandler{DB: pgdb}
	r.Get("/", bookHandler.ListBooks)
	r.Get("/{id}", bookHandler.GetBook)
	r.Post("/", bookHandler.CreateBook)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)

	return r
}
