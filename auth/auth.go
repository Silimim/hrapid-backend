package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type loginReqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
		if tokenString == "" {
			http.Error(w, "missing authorization token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("HRAPID_SECRET")), nil
		})

		if err != nil {
			println(err.Error())
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := int32(claims["sub"].(float64))
			var user model.User
			result := db.GetDB().Where("id = ?", userID).First(&user)
			if result.Error != nil {
				http.Error(w, "user not found", http.StatusNotFound)
				return
			}

			type contextKey string

			const userKey contextKey = "user"

			ctx := context.WithValue(r.Context(), userKey, userID)
			r = r.WithContext(ctx)
		} else {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Register(w http.ResponseWriter, r *http.Request) {

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

	err = json.NewEncoder(w).Encode("user created")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user model.User
	var loginReqBody loginReqBody

	json.NewDecoder(r.Body).Decode(&loginReqBody)

	if loginReqBody.Username == "" || loginReqBody.Password == "" {
		http.Error(w, "username and password cannot be empty", http.StatusBadRequest)
		return
	}

	result := db.GetDB().Where("username = ?", loginReqBody.Username).First(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		http.Error(w, "error during user search", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReqBody.Password)); err != nil {
		http.Error(w, "invalid username or password", http.StatusBadRequest)
		return
	}

	token, err := generateTokenPair(user)
	if err != nil {
		http.Error(w, "error during token generation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Refresh(w http.ResponseWriter, r *http.Request) {

	tokenString := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	if tokenString == "" {
		http.Error(w, "missing authorization token", http.StatusUnauthorized)
		return
	}

	var user model.User

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("HRAPID_SECRET")), nil
	})

	if err != nil {
		println(err.Error())
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int32(claims["sub"].(float64))
		db.GetDB().Where("id = ?", userID).First(&user)

		if userID == user.ID {

			newTokenPair, err := generateTokenPair(user)
			if err != nil {
				http.Error(w, "error in token generation", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)

			err = json.NewEncoder(w).Encode(newTokenPair)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "invalid user", http.StatusBadRequest)
			return
		}

	} else {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
}
