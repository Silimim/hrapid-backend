package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Silimim/hrapid-backend/api"
	"github.com/Silimim/hrapid-backend/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.InitDB()

	_, err = db.INITORM()
	if err != nil {
		log.Fatalf("Error starting GORM")
	} else {
		log.Printf("GORM started successfully")
	}

	router := mux.NewRouter()
	router.HandleFunc("/health", HealthCheckHandler)
	router.HandleFunc("/users", api.GetUsers).Methods("GET")
	router.HandleFunc("/users", api.AddUser).Methods("POST")

	log.Printf("App starting on port %s", os.Getenv("HRAPID_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HRAPID_PORT"), router))

}
