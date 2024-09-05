package api

import (
	"encoding/json"
	"net/http"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	var employee []*model.Employee

	db.GetDB().Find(&employee)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	var employee *model.Employee

	db.GetDB().Find(&employee, r.URL.Query().Get("id"))
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var employee model.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.GetDB().Create(&employee)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
