package server

import (
	"context"
	"fmt"
	"go-api/internal/books"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg"
	"github.com/markbates/goth/gothic"
)

type ContextKey string

const providerKey ContextKey = "provider"

func ConfigureRoutes(pgdb *pg.DB) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("We're up and running!")))
	})

	r.Get("/auth/{provider}/callback", getAuthCallbackFunc)
	r.Get("/auth/{provider}", beginAuthCallback)
	r.Get("/logout/{provider}", logout)

	r.Mount("/books", books.Routes(pgdb))

	return r

}

func getAuthCallbackFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You hit the auth callback!")

	// context gymnastics because of goth interaction with chi...
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), providerKey, provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(user)

	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}

func beginAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), providerKey, provider))
	gothic.BeginAuthHandler(w, r)
}

func logout(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), providerKey, provider))
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
