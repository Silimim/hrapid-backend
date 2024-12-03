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

	if company.ID == 0 {
		http.Error(w, "Company ID is required", http.StatusBadRequest)
		return
	}

	result := db.GetDB().Model(&company).Updates(&company)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Company updated successfully")
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {

	var company model.Company

	result := db.GetDB().Find(&company, r.URL.Query().Get("id"))
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	db.GetDB().Delete(&company)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("Company deleted successfully")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
