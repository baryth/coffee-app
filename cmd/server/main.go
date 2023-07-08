package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Printf("API listening on port: %s", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// TODO: add router
	}

	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	// TODO: connection to the db

	app := &Application{
		Config: cfg,
		// TODO: add models later
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
