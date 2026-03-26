package handlers

import (
	"net/http"

	"github.com/aprimr/chautari/services"
	"github.com/aprimr/chautari/utils"
	"github.com/aprimr/chautari/validation"
	"github.com/go-chi/chi/v5"
)

func SendContactRequestHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// get reciever_id param from url
	receiverId := chi.URLParam(r, "reciever_id")
	if validation.IsEmptyString(receiverId) {
		utils.SendError(w, "Empty contact id", http.StatusBadRequest)
		return
	}
	if receiverId == uid {
		utils.SendError(w, "Cannot send request", http.StatusConflict)
		return
	}

	// Call Add Contact service
	err := services.SendContactRequest(r.Context(), uid, receiverId)
	if err != nil {
		if err.Error() == "request exists" {
			utils.SendError(w, "Request already sent", http.StatusConflict)
			return
		}
		utils.LogError("SendContactRequest service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "R			equest sent", nil, http.StatusCreated)
}
