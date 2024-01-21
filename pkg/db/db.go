package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

func StartDB() (*pg.DB, error) {
	var (
		opts *pg.Options
		err  error
	)

	if os.Getenv("ENV") == "PROD" {
		opts, err = pg.ParseURL(os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, err
		}
	} else {
		opts = &pg.Options{

			Addr:     "db:5432",
			User:     "myuser",
			Password: "admin",
			Database: "mydb"}
	}

	//connect db
	db := pg.Connect(opts)
	if db == nil {
		log.Println("Failed to connect to database")
		return nil, fmt.Errorf("failed to connect to database")
	}

	//run migrations
	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		log.Printf("Error discovering migrations: %v\n", err)
		return nil, err
	}

	//start the migrations
	_, _, err = collection.Run(db, "init")
	if err != nil {
		log.Printf("Error during migration init: %v\n", err)
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		log.Printf("Error running migrations: %v\n", err)
		return nil, err
	}
	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Printf("version is %d\n", oldVersion)
	}

	//return the db connection
	return db, err
}
