package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
)

func GetCompanies(w http.ResponseWriter, r *http.Request) {

	var companies []*model.Company

	db.GetDB().Find(&companies)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(companies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCompany(w http.ResponseWriter, r *http.Request) {

	var company *model.Company

	db.GetDB().Find(&company, r.URL.Query().Get("id"))
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var company model.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		log.Printf("Error decoding company: %v", err)
		http.Error(w, "Invalid request body format", http.StatusBadRequest)
		return
	}

	date := time.Now()
	company.DateAdded = &date

	type contextKey string
	const userKey contextKey = "user"

	userValue := r.Context().Value(userKey)
	if userValue == nil {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	userID, ok := userValue.(int32)
	if !ok {
		log.Printf("Invalid user type in context. Got: %T, want: int32", userValue)
		http.Error(w, "Invalid user ID type", http.StatusInternalServerError)
		return
	}

	company.UserAddedID = &userID

	result := db.GetDB().Create(&company)
	if result.Error != nil {
		log.Printf("Error creating company: %v", result.Error)
		http.Error(w, "Error creating company", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Company created successfully",
		"id":      company.ID,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {

	var company model.Company

	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.GetDB().Save(&company)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode("Company updated successfully")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {

	var company model.Company

	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.GetDB().Delete(&company)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode("Company deleted successfully")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
