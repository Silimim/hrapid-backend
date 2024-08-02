package api

import (
	"encoding/json"
	"net/http"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []*model.User

	db.GetDB().Find(&users)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	var user *model.User

	db.GetDB().Find(&user, r.URL.Query().Get("id"))
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
