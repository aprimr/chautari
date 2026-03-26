package handlers

import (
	"net/http"

	"github.com/aprimr/chautari/services"
	"github.com/aprimr/chautari/utils"
	"github.com/aprimr/chautari/validation"
)

func GetMeHandler(w http.ResponseWriter, r *http.Request) {
	// Get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Call GetMe service
	userData, err := services.GetMe(r.Context(), uid)
	if err != nil {
		if err.Error() == "user is inactive" {
			utils.SendError(w, "Your account has been deactivated", http.StatusForbidden)
			return
		}
		if err.Error() == "user is deleted" {
			utils.SendError(w, "Your account no longer exists", http.StatusForbidden)
			return
		}
		utils.LogError("Get Me", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Fecthed successfully", userData, http.StatusOK)
}

func SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// get search query from URL
	search := r.URL.Query().Get("q")
	if validation.IsEmptyString(search) {
		utils.SendError(w, "Search query cannot be empty", http.StatusBadRequest)
		return
	}
	if len(search) < 3 {
		utils.SendError(w, "Search query must be at least 3 characters", http.StatusBadRequest)
		return
	}

	// Call service
	users, err := services.SearchUser(r.Context(), search, uid)
	if err != nil {
		utils.LogError("SearchUser service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Search successful", users, http.StatusOK)
}
