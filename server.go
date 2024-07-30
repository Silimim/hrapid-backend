package main

import (
	"log"
	"net/http"

	"github.com/Silimim/hrapid-backend/api"
	"github.com/Silimim/hrapid-backend/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var database *gorm.DB

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDB()

	database, err = db.INITORM()
	if err != nil {
		log.Fatalf("Error starting GORM")
	} else {
		log.Printf("GORM started successfully")
	}

	router := mux.NewRouter()
	router.HandleFunc("/users", api.GetUsers).Methods("GET")
	router.HandleFunc("/users", api.AddUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
