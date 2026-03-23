package main

import (
	"net/http"
	"os"

	"github.com/aprimr/chautari/db"
	"github.com/aprimr/chautari/handlers"
	"github.com/aprimr/chautari/utils"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		utils.LogFatal("Error loading env file", err)
	}

	// Create router
	r := chi.NewRouter()

	// Connect to DB
	db.ConnectDB()

	// Routes
	r.Route("chautari/api/v1", func(r chi.Router) {

		r.Post("/register", handlers.UserRegistrationHandler)

	})

	// Spin Up server
	port := ":" + os.Getenv("PORT")
	utils.LogInfo("Server started on port" + port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		utils.LogFatal("Error starting server: ", err)
	}
}
