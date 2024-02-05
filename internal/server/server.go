package server

import (
	"fmt"
	db "go-api/internal/storage"
	"log"
	"net/http"
	"os"
)

func Startup() {

	log.Print("server has started up...")

	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error starting the database %v", err)
	}

	router := ConfigureRoutes(pgdb)

	port := os.Getenv("PORT")

	err = http.ListenAndServe((fmt.Sprintf(":%s", port)), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
