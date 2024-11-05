package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Silimim/hrapid-backend/api"
	"github.com/Silimim/hrapid-backend/auth"
	"github.com/Silimim/hrapid-backend/db"
	"github.com/gorilla/handlers"
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

	db.InitORM()

	if err != nil {
		log.Fatalf("Error starting GORM")
	} else {
		log.Printf("GORM started successfully")
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE", "PATCH", "UPDATE"})

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	router.HandleFunc("/health", HealthCheckHandler)
	router.HandleFunc("/users", api.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	router.HandleFunc("/companies", api.GetCompanies).Methods("GET")
	router.HandleFunc("/companies/{id}", api.GetCompany).Methods("GET")
	router.HandleFunc("/companies", api.CreateCompany).Methods("POST")

	router.HandleFunc("/register", auth.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")

	log.Printf("App starting on port %s", os.Getenv("HRAPID_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HRAPID_PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
