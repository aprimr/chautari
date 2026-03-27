package handlers

import (
	"net/http"

	"github.com/aprimr/chautari/services"
	"github.com/aprimr/chautari/utils"
	"github.com/aprimr/chautari/validation"
	"github.com/go-chi/chi/v5"
)

func GetFriendsHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Call service
	friends, err := services.GetFriends(r.Context(), uid)
	if err != nil {
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Friends fetched", friends, http.StatusOK)
}

func SendRequestHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// get reciever_id param from url
	receiverId := chi.URLParam(r, "receiver_id")
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

	// Call service
	err := services.AcceptRequest(r.Context(), requestId, uid)
	if err != nil {
		utils.LogError("AcceptRequest service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "You are now friends", nil, http.StatusOK)
}

func RejectRequestHandler(w http.ResponseWriter, r *http.Request) {
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

	// Call service
	err := services.RejectRequest(r.Context(), requestId, uid)
	if err != nil {
		utils.LogError("RejectRequest service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Request rejected", nil, http.StatusOK)
}

func UnfriendHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// get rid from url params
	rid := chi.URLParam(r, "rid")
	if validation.IsEmptyString(rid) {
		utils.SendError(w, "Empty rid", http.StatusBadRequest)
		return
	}

	// Call service
	err := services.UnfriendUser(r.Context(), rid, uid)
	if err != nil {
		if err.Error() == "user not found" {
			utils.SendError(w, "User not found", http.StatusNotFound)
			return
		}
		utils.LogError("UnfriendUser service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Unfriend successful", nil, http.StatusOK)
}

func GetIncomingRequestsHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Call service
	requests, err := services.GetIncomingRequests(r.Context(), uid)
	if err != nil {
		utils.LogError("GetIncomingRequests service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Incoming requests fetched", requests, http.StatusOK)
}

func GetOutgoingRequestsHandler(w http.ResponseWriter, r *http.Request) {
	// get uid from r.context
	uid, ok := r.Context().Value("uid").(string)
	if !ok || uid == "" {
		utils.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Call service
	requests, err := services.GetOutgoingRequests(r.Context(), uid)
	if err != nil {
		utils.LogError("GetOutgoingRequests service", err)
		utils.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Outgoing requests fetched", requests, http.StatusOK)
}
