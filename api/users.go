package api

import (
	"encoding/json"
	"net/http"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
)

type UserData struct {
	ID       int32   `json:"id"`
	Name     string  `json:"name"`
	LastName *string `json:"last_name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Phone    *string `json:"phone"`
	Role     string  `json:"role"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []*model.User
	var userData []*UserData

	db.GetDB().Find(&users)

	for _, user := range users {
		userData = append(userData, &UserData{
			ID:       user.ID,
			Name:     user.Name,
			LastName: user.LastName,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			Role:     user.Role,
		})
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(userData)
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

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {

	var user *model.User

	db.GetDB().Find(&user, r.URL.Query().Get("username"))
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "payload error", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		http.Error(w, "username, password and email cannot be empty", http.StatusBadRequest)
		return
	}

	if err := db.GetDB().Where("username = ? OR email = ?", user.Username, user.Email).First(&user).Error; err == nil {
		http.Error(w, "username or email already exists", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(user.Email) {
		http.Error(w, "email is not in a valid format", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "error during password hashing", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	db.GetDB().Create(&user)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
