package main

import (
	"net/http"
	"os"

	"github.com/aprimr/chautari/db"
	"github.com/aprimr/chautari/handlers"
	"github.com/aprimr/chautari/middlewares"
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
	r.Route("/chautari/api/v1", func(r chi.Router) {

		// Public Routes

		// Auth
		r.Post("/login", handlers.UserLoginHandler)
		r.Post("/register", handlers.UserRegistrationHandler)

		// Protected Routes (user)
		r.Group(func(r chi.Router) {
			r.Use(middlewares.Authentication)

			r.Get("/me", handlers.GetMeHandler)

			// Search User
			r.Get("/search", handlers.SearchUserHandler)

			// Contact routes
			r.Route("/requests", func(r chi.Router) {
				r.Post("/send/{receiver_id}", handlers.SendRequestHandler) // send request to contact_id
				r.Delete("/cancel/{rid}", handlers.CancelRequestHandler)   //Cancle request
				r.Post("/accept/{rid}", handlers.AcceptRequestHandler)     // accept request
				r.Post("/reject/{rid}", handlers.RejectRequestHandler)     // reject request
			})
		})

	})

	// Spin Up server
	port := ":" + os.Getenv("PORT")
	utils.LogInfo("Server started on port" + port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		utils.LogFatal("Error starting server: ", err)
	}
}
