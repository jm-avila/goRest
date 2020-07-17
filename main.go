package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo/handlers"
	"todo/postgres"

	"github.com/go-pg/pg"
)

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "admin",
		Database: "todo_dev",
	})

	r := handlers.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	defer DB.Close()

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)

	if err != nil {
		log.Fatalf("Cannot start server %v", err)
	}
}
