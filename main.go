package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo/handlers"
	"todo/postgres"
	"todo/domain"

	"github.com/go-pg/pg/v9"
)

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "admin",
		Database: "todo_dev",
	})

	defer DB.Close()

	domainDB := domain.DB{
		UserRepo: postgres.NewUserRepo(DB),
	}

	d := &domain.Domain{DB: domainDB}

	r := handlers.SetupRouter(d)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)

	if err != nil {
		log.Fatalf("Cannot start server %v", err)
	}
}
