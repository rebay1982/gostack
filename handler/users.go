package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/rebay1982/gostack/db"
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
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

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

	userDb, err := db.GetUserById(id)

	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read from DB: %v", err), http.StatusInternalServerError)
		return
	}

	// If no error, make sure we got a result.
	if userDb == nil {
		http.Error(w, fmt.Sprintf("No user with id [%d]", id), http.StatusNotFound)
		return

	} else {
		output := fmt.Sprintf("userGet: %v", *userDb)
		w.Write([]byte(output))
	}
}
