package main

import (
	"backend/internal/repository"
	dprepo "backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
	DSN    string
	DB     repository.DatabaseRepo
}

func main() {
	// set application config
	var app application
	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=root password=secret dbname=assistant_teacher sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()
	// connect to the db
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dprepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.Domain = "example.com"

	log.Println("Starting application on port ", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
