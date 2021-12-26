package main

import (
	"log"
	"net/http"
	// "go-library/models"
	// "go-library/services"
	"go-library/routers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const username, plainTextPassword, fullName, email = "suresh", "password", "Suresh Baskar", "suresh@test.com"

func main() {
	godotenv.Load(".env")
	server := mux.NewRouter()
	server.HandleFunc("/users/create", routers.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", server))
}
