package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aprimr/chautari/models"
	"github.com/aprimr/chautari/services"
	"github.com/aprimr/chautari/utils"
	"github.com/aprimr/chautari/validation"
)

func UserRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON and store in registerInput
	registerInput := models.RegisterInput{}
	err := json.NewDecoder(r.Body).Decode(&registerInput)
	if err != nil {
		utils.LogError("Invalid JSON", err)
		utils.SendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate user data
	if validation.IsEmptyString(registerInput.Name) {
		utils.SendError(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}
	if !validation.IsValidEmail(registerInput.Email) {
		utils.SendError(w, "Invalid email", http.StatusBadRequest)
		return
	}
	if msg, isValid := validation.IsValidPassword(registerInput.Password); !isValid {
		utils.SendError(w, msg, http.StatusBadRequest)
		return
	}

	// Call RegisterUser service
	err = services.RegisterUser(r.Context(), registerInput)
	if err != nil {
		if err.Error() == "email already exists" {
			utils.LogError("User already exists with this email", err)
			utils.SendError(w, "Email already in use", http.StatusConflict)
			return
		}
		if err.Error() == "could not generate unique username" {
			utils.LogError("Could not generate unique username", err)
			utils.SendError(w, "Error generating unique username", http.StatusConflict)
			return
		}
		utils.LogError("Register Service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "User registered successfully", nil, http.StatusCreated)
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON and store in registerInput
	loginInput := models.LoginInput{}
	err := json.NewDecoder(r.Body).Decode(&loginInput)
	if err != nil {
		utils.LogError("Invalid JSON", err)
		utils.SendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate user data
	if !validation.IsValidEmail(loginInput.Email) {
		utils.SendError(w, "Invalid email", http.StatusBadRequest)
		return
	}
	if msg, isValid := validation.IsValidPassword(loginInput.Password); !isValid {
		utils.SendError(w, msg, http.StatusBadRequest)
		return
	}

	// Call LoginUser service
	jwtToken, err := services.LoginUser(r.Context(), loginInput)
	if err != nil {
		if err.Error() == "user not found" {
			utils.LogError("User not found", err)
			utils.SendError(w, "User not found", http.StatusNotFound)
			return
		}
		if err.Error() == "incorrect password" {
			utils.LogError("Incorrect password", err)
			utils.SendError(w, "Incorrect password", http.StatusUnauthorized)
			return
		}
		utils.LogError("Login Service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "User logged in successfully", jwtToken, http.StatusOK)
}
