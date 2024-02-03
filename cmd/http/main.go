package main

import (
	"fmt"
	db "go-api/internal/storage"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg"

	"go-api/internal/books"
)

func main() {
	log.Print("server has started up...")

	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error starting the database %v", err)
	}

	router := startAPI(pgdb)

	port := os.Getenv("PORT")

	err = http.ListenAndServe((fmt.Sprintf(":%s", port)), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}

func startAPI(pgdb *pg.DB) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("We're up and running!")))
	})
	// passing the db connection explicitly to each handler
	r.Mount("/books", books.Routes(pgdb))

	return r

}
