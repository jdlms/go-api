package db

import (
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
	//run migrations
	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		return nil, err
	}

	//start the migrations
	_, _, err = collection.Run(db, "init")
	if err != nil {
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
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
