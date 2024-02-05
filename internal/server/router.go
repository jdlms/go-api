package server

import (
	"fmt"
	"go-api/internal/books"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg"
	"github.com/markbates/goth/gothic"
)

func ConfigureRoutes(pgdb *pg.DB) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// r.Get("/auth/{provider}/callback", s.getAuthCallbackFunc)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("We're up and running!")))
	})
	// passing the db connection explicitly to each handler

	r.Get("/auth/{provider}", getAuthCallbackFunc)

	r.Mount("/books", books.Routes(pgdb))

	return r

}

func getAuthCallbackFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You hit the auth callback!")
	// provider := chi.URLParam(r, "provider")
	// r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(user)
}
