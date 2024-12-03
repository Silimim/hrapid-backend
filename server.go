package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Silimim/hrapid-backend/api"
	"github.com/Silimim/hrapid-backend/auth"
	"github.com/Silimim/hrapid-backend/db"
	"github.com/Silimim/hrapid-backend/table"
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
	apiRouter.HandleFunc("/user/{id}", api.GetUser).Methods("GET")
	apiRouter.HandleFunc("/user", api.CreateUser).Methods("POST")
	apiRouter.HandleFunc("/user", api.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/user/{id}", api.DeleteUser).Methods("DELETE")

	apiRouter.HandleFunc("/companies", api.GetCompanies).Methods("GET")
	apiRouter.HandleFunc("/company/{id}", api.GetCompany).Methods("GET")
	apiRouter.HandleFunc("/company", api.CreateCompany).Methods("POST")
	apiRouter.HandleFunc("/company", api.UpdateCompany).Methods("PUT")
	apiRouter.HandleFunc("/company/{id}", api.DeleteCompany).Methods("DELETE")

	apiRouter.HandleFunc("/employees", api.GetEmployees).Methods("GET")
	apiRouter.HandleFunc("/employee/{id}", api.GetEmployee).Methods("GET")
	apiRouter.HandleFunc("/employee", api.CreateEmployee).Methods("POST")
	apiRouter.HandleFunc("/employee", api.UpdateEmployee).Methods("PUT")
	apiRouter.HandleFunc("/employee/{id}", api.DeleteEmployee).Methods("DELETE")

	apiRouter.HandleFunc("/lists", api.GetLists).Methods("GET")
	apiRouter.HandleFunc("/list/{id}", api.GetList).Methods("GET")
	apiRouter.HandleFunc("/list", api.CreateList).Methods("POST")
	apiRouter.HandleFunc("/list", api.UpdateList).Methods("PUT")
	apiRouter.HandleFunc("/list/{id}", api.DeleteList).Methods("DELETE")

	tableRouter := r.PathPrefix("/table").Subrouter()
	tableRouter.Use(auth.JwtAuthentication)
	tableRouter.HandleFunc("/companies", table.CompaniesHandler).Methods("GET")
	tableRouter.HandleFunc("/employees", table.EmployeesHandler).Methods("GET")
	tableRouter.HandleFunc("/lists", table.ListsHandler).Methods("GET")

	log.Printf("App starting on port %s", os.Getenv("HRAPID_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("HRAPID_PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
