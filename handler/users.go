package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rebay1982/gostack/models"
	"github.com/rebay1982/gostack/server"
)

// Users Handler for users.
func Users(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		userGet(w, r)
		return

	case http.MethodPost:
		userPost(w, r)
		return

	case http.MethodDelete:
		userDelete(w, r)
		return

	default:
		w.Header().Set("Allow", http.MethodGet)
		w.Header().Add("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

func userGet(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	// Validate parameter is an int.
	if err != nil {

		http.Error(w, fmt.Sprintf("Invalid 'id' Parameter: %s", idStr), http.StatusBadRequest)
		return
	}

	user, err := server.GetUserById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, fmt.Sprintf("No user with id [%d]", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*user)
}

func userPost(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		errorMsg := fmt.Sprintf("Cannot parse JSON request body: %v", err)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	err = server.CreateUser(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func userDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	// Validate parameter is an int.
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid 'id' Parameter: %s", idStr), http.StatusBadRequest)
		return
	}

	err = server.DeleteUserById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}
}
