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

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	var emplyees []*model.Employee

	db.GetDB().Find(&emplyees)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(emplyees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	var emplyee *model.Employee

	db.GetDB().Find(&emplyee, r.URL.Query().Get("id"))
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(emplyee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var emplyee model.Employee
	err := json.NewDecoder(r.Body).Decode(&emplyee)
	if err != nil {
		log.Printf("Error decoding emplyee: %v", err)
		http.Error(w, "Invalid request body format", http.StatusBadRequest)
		return
	}

	date := time.Now()
	emplyee.DateAdded = &date

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

	userID := user.ID

	emplyee.UserAddedID = &userID

	result := db.GetDB().Create(&emplyee)
	if result.Error != nil {
		log.Printf("Error creating emplyee: %v", result.Error)
		http.Error(w, "Error creating emplyee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Employee created successfully",
		"id":      emplyee.ID,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var emplyee model.Employee

	err := json.NewDecoder(r.Body).Decode(&emplyee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if emplyee.ID == 0 {
		http.Error(w, "Company ID is required", http.StatusBadRequest)
		return
	}

	result := db.GetDB().Model(&emplyee).Updates(&emplyee)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Employee updated successfully")
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	var employee model.Employee

	result := db.GetDB().Find(&employee, r.URL.Query().Get("id"))
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	db.GetDB().Delete(&employee)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("Employee deleted successfully")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
