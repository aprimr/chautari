package handlers

import (
	"net/http"

	"github.com/aprimr/chautari/services"
	"github.com/aprimr/chautari/utils"
	"github.com/aprimr/chautari/validation"
	"github.com/go-chi/chi/v5"
)

func SendRequestHandler(w http.ResponseWriter, r *http.Request) {
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
	err := services.SendRequest(r.Context(), uid, receiverId)
	if err != nil {
		if err.Error() == "request exists" {
			utils.SendError(w, "Request already sent", http.StatusConflict)
			return
		}
		utils.LogError("SendContactRequest service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Request sent", nil, http.StatusCreated)
}

func CancelRequestHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// get rid from url params
	rid := chi.URLParam(r, "rid")
	if validation.IsEmptyString(rid) {
		utils.SendError(w, "Empty request id", http.StatusBadRequest)
		return
	}

	// Call service
	err := services.CancelRequest(r.Context(), rid, uid)
	if err != nil {
		if err.Error() == "request not found" {
			utils.SendError(w, "Request not found", http.StatusNotFound)
			return
		}
		utils.LogError("CancleRequest service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Request cancelled", nil, http.StatusOK)
}

func AcceptRequestHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {

		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// get id from url params (request id)
	requestId := chi.URLParam(r, "rid")
	if validation.IsEmptyString(requestId) {
		utils.SendError(w, "Empty request id", http.StatusBadRequest)
		return
	}

	utils.LogDebug("Rid handler: " + requestId)

	// Call service
	err := services.AcceptRequest(r.Context(), requestId, uid)
	if err != nil {
		utils.LogError("AcceptRequest service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "You are now friends", nil, http.StatusOK)
}
