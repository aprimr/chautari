package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aprimr/chautari/db"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading env file", err)
	}

	// Create router
	r := chi.NewRouter()

	// Connect to DB
	db.ConnectDB()

	// Spin Up server
	port := ":" + os.Getenv("PORT")
	log.Println("Server started on port", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}
