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

func Register(w http.ResponseWriter, r *http.Request) (string, error) {

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", err
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		return "", errors.New("username, password, and email cannot be empty")
	}

	if err := db.GetDB().Where("username = ? OR email = ?", user.Username, user.Email).First(&user).Error; err == nil {
		return "", errors.New("username or email already exists")
	}

	if !isValidEmail(user.Email) {
		return "", errors.New("invalid email format")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword

	db.GetDB().Create(&user)

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(w http.ResponseWriter, r *http.Request) (string, error) {

	var user model.User

	if user.Username == "" || user.Password == "" {
		return "", errors.New("username and password cannot be empty")
	}

	result := db.GetDB().Where("username = ?", user.Username).First(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
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
