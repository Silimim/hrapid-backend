package api

import (
	"encoding/json"
	"net/http"

	"github.com/Silimim/hrapid-backend/db/query"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := query.User.Find()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user = query.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = user.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
