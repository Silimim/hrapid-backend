package api

import (
	"encoding/json"
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

	var company model.Company

	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date := time.Now()
	company.DateAdded = &date

	print(r.Context().Value("user"))

	if r.Context().Value("user") == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user := r.Context().Value("user")
	userID := (user.(model.User)).ID

	company.UserAddedID = &userID

	db.GetDB().Create(&company)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode("Company created successfully")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
