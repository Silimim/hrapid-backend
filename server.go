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

	r := mux.NewRouter()

	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/refresh", auth.Refresh).Methods("POST")

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(auth.JwtAuthentication)
	apiRouter.HandleFunc("/health", HealthCheckHandler)

	apiRouter.HandleFunc("/users", api.GetUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users/username/{username}", api.GetUserByUsername).Methods("GET")
	apiRouter.HandleFunc("/users", api.CreateUser).Methods("POST")

	apiRouter.HandleFunc("/companies", api.GetCompanies).Methods("GET")
	apiRouter.HandleFunc("/companies/{id}", api.GetCompany).Methods("GET")
	apiRouter.HandleFunc("/companies", api.CreateCompany).Methods("POST")

	apiRouter.HandleFunc("/employees", api.GetEmployees).Methods("GET")
	apiRouter.HandleFunc("/employees/{id}", api.GetEmployee).Methods("GET")
	apiRouter.HandleFunc("/employees", api.CreateEmployee).Methods("POST")

	apiRouter.HandleFunc("/lists", api.GetLists).Methods("GET")
	apiRouter.HandleFunc("/lists/{id}", api.GetList).Methods("GET")
	apiRouter.HandleFunc("/lists", api.CreateList).Methods("POST")

	log.Printf("App starting on port %s", os.Getenv("HRAPID_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HRAPID_PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
