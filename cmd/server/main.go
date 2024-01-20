package main

import (
	"fmt"
	"go-api/pkg/api"
	"go-api/pkg/db"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("server has started up...")

	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error starting the database %v", err)
	}

	router := api.StartAPI(pgdb)

	port := os.Getenv("PORT")

	err = http.ListenAndServe((fmt.Sprintf(":%s", port)), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
