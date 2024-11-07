package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/db/model"
	"github.com/Silimim/hrapid-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "missing authorization token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("secret"), nil
		})

		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := int(claims["sub"].(float64))
			var user model.User
			result := db.GetDB().Where("id = ?", userID).First(&user)
			if result.Error != nil {
				http.Error(w, "user not found", http.StatusNotFound)
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)
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
		return
	}

	result := db.GetDB().Where("username = ?", user.Username).First(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
		}
		http.Error(w, "uerror during user search", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		http.Error(w, "invalid username or password", http.StatusBadRequest)
		return
	}

	token, err := generateTokenPair(user)
	if err != nil {
		http.Error(w, "uerror during token generation", http.StatusInternalServerError)
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

	var tokenReq tokenReqBody

	err := json.NewDecoder(r.Body).Decode(&tokenReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user *model.User

	token, _ := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		db.GetDB().Where("id = ?", int(claims["sub"].(float64))).First(&user)
		if claims["sub"] == user.ID {

			newTokenPair, err := generateTokenPair(*user)
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
		}

		http.Error(w, "invalid user", http.StatusBadRequest)
		return
	}
}
