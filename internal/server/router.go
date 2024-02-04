package server

import (
	"go-api/internal/auth"
	"go-api/internal/books"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg"
)

func ConfigureRoutes(pgdb *pg.DB) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// r.Get("/auth/{provider}/callback", s.getAuthCallbackFunc)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("We're up and running!")))
	})
	// passing the db connection explicitly to each handler
	r.Mount("/", auth.Routes())
	r.Mount("/books", books.Routes(pgdb))

	return r

}
