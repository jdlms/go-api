package auth

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/auth/github/callback", getAuthCallbackFunc)

	return r
}

func getAuthCallbackFunc(w http.ResponseWriter, r *http.Request) {

	// provider := chi.URLParam(r, "provider")

	// r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(user)
}
