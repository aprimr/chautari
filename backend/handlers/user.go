package handlers

import (
	"net/http"

	"github.com/aprimr/chautari/services"
	"github.com/aprimr/chautari/utils"
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
