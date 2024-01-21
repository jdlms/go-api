package api

import (
	"net/http"

	"go-api/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg"
)

func StartAPI(pgdb *pg.DB) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	// what is middleware with value?

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("We're up and running!")))
	})
	// passing the db connection explicitly to each handler
	r.Mount("/books", BookRoutes(pgdb))

	// http.ListenAndServe(":8080", r)

	return r

}

func BookRoutes(pgdb *pg.DB) chi.Router {
	r := chi.NewRouter()

	bookHandler := handlers.BookHandler{DB: pgdb}
	r.Get("/", bookHandler.ListBooks)
	// r.Get("/{id}", bookHandler.GetBook)
	// r.Post("/", bookHandler.CreateBook)
	// r.Put("/{id}", bookHandler.UpdateBook)
	// r.Delete("/{id}", bookHandler.DeleteBook)

	return r
}
