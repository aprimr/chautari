package handlers

import (
	"net/http"

	"github.com/aprimr/chautari/services"
	"github.com/aprimr/chautari/utils"
	"github.com/aprimr/chautari/validation"
)

func SendContactRequestHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// get contact_id query from url
	contactId := r.URL.Query().Get("contact_id")
	if validation.IsEmptyString(contactId) {
		utils.SendError(w, "empty contact_id", http.StatusBadRequest)
		return
	}
	if contactId == uid {
		utils.SendError(w, "Cannot add self to contact", http.StatusConflict)
		return
	}

	// Call Add Contact service
	err := services.SendContactRequest(r.Context(), uid, contactId)
	if err != nil {
		if err.Error() == "request exists" {
			utils.SendError(w, "Request already sent", http.StatusConflict)
			return
		}
		utils.LogError("SendContactRequest service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Contact request sent", nil, http.StatusCreated)
}
