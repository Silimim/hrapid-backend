package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []*model.User

	db.GetDB().Find(&users)

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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding user: %v", err)
		http.Error(w, "Invalid request body format", http.StatusBadRequest)
		return
	}

	userValue := r.Context().Value(utils.UserKey)
	if userValue == nil {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	user, ok := userValue.(model.User)
	if !ok {
		http.Error(w, "Invalid user type in context", http.StatusInternalServerError)
		return
	}

	result := db.GetDB().Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "User created successfully",
		"id":      user.ID,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	result := db.GetDB().Model(&user).Updates(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User updated successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	result := db.GetDB().Find(&user, r.URL.Query().Get("id"))
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	db.GetDB().Delete(&user)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("User deleted successfully")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
