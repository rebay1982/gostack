package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rebay1982/gostack/models"
)

// Users Handler for users.
func Users(w http.ResponseWriter, r *http.Request) {

	// do POST and GET methods only.
	if r.Method == http.MethodGet {
		userGet(w, r)

		return

	} else if r.Method == http.MethodPost {
		userPost(w, r)

		return

	} else {
		w.Header().Set("Allow", http.MethodGet)
		w.Header().Add("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)

		return
	}

	return
}

func userPost(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		errorMsg := fmt.Sprintf("Cannot parse JSON request body: %v", err)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	output := fmt.Sprintf("userPost: %v", user)
	w.Write([]byte(output))
}

func userGet(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	// Validate parameter is an int.
	if err != nil {

		http.Error(w, fmt.Sprintf("Invalid 'id' Parameter: %s", idStr), http.StatusBadRequest)
		return
	}

	output := fmt.Sprintf("userGet: %d", id)
	w.Write([]byte(output))
}
