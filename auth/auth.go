package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "payload error", http.StatusBadRequest)
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		http.Error(w, "username, password and email cannot be empty", http.StatusBadRequest)
	}

	if err := db.GetDB().Where("username = ? OR email = ?", user.Username, user.Email).First(&user).Error; err == nil {
		http.Error(w, "username or email already exists", http.StatusBadRequest)
	}

	if !isValidEmail(user.Email) {
		http.Error(w, "email is not in a valid format", http.StatusBadRequest)
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "error during password hashing", http.StatusInternalServerError)
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

func Login(w http.ResponseWriter, r *http.Request) {

	var user model.User

	if user.Username == "" || user.Password == "" {
		http.Error(w, "username and password cannot be empty", http.StatusBadRequest)
	}

	result := db.GetDB().Where("username = ?", user.Username).First(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
		}
		http.Error(w, "uerror during user search", http.StatusInternalServerError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		http.Error(w, "invalid username or password", http.StatusBadRequest)
	}

	token, err := generateJWT(user)
	if err != nil {
		http.Error(w, "uerror during token generation", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func generateJWT(user model.User) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
