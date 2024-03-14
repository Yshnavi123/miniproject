// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Model Struct
type User struct {
	ID        uint
	Name      string
	Age       int
	Birthday  time.Time
	CreatedAt time.Time
}

var db *gorm.DB
var err error

func main() {
	time.Sleep(10 * time.Second)
	// Connect to PostgreSQL Database

	db, err = gorm.Open(postgres.Open("postgresql://postgres:postgres@db-service:5432/progressdb?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	// AutoMigrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	// Initialize HTTP server with Gorilla Mux
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", GetUserByIDHandler).Methods("GET")

	// Start HTTP server
	fmt.Println("HTTP server listening on port 8080")
	http.ListenAndServe(":8080", r)
}

// HTTP Handler functions
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create user using HTTP
	// Example:
	
	decoder := json.NewDecoder(r.Body)
    var u User
    err := decoder.Decode(&u)
	if err != nil{
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	
	result := db.Create(&u)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully!")
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve user by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	userID := vars["id"]
	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Return user details as JSON response
	json.NewEncoder(w).Encode(user)
}
