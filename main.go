package main

import (
	"log"
	"os"

	"./handlers"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {
	var err error
	username := os.Getenv("username")
	db, err = sqlx.Connect("postgres", "user="+username+" dbname=yourdbname sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()

	userHandler := handlers.NewUserHandler(userModel)
	r.Post("/register", userHandler.Register)
}
