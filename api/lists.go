package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
)

func GetLists(w http.ResponseWriter, r *http.Request) {

	var lists []*model.List

	db.GetDB().Find(&lists)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(lists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetList(w http.ResponseWriter, r *http.Request) {

	var list *model.List

	db.GetDB().Find(&list, r.URL.Query().Get("id"))
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var list model.List
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		log.Printf("Error decoding list: %v", err)
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

	date := time.Now()
	list.DateAdded = &date

	userID := user.ID

	list.UserAddedID = &userID

	result := db.GetDB().Create(&list)
	if result.Error != nil {
		log.Printf("Error creating list: %v", result.Error)
		http.Error(w, "Error creating list", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "List created successfully",
		"id":      list.ID,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func UpdateList(w http.ResponseWriter, r *http.Request) {
	var list model.List

	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if list.ID == 0 {
		http.Error(w, "List ID is required", http.StatusBadRequest)
		return
	}

	result := db.GetDB().Model(&list).Updates(&list)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "List not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("List updated successfully")
}

func DeleteList(w http.ResponseWriter, r *http.Request) {

	var list model.List

	result := db.GetDB().Find(&list, r.URL.Query().Get("id"))
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	db.GetDB().Delete(&list)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("List deleted successfully")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
