// package main

import (
	"api-golang/handlers"
	"api-golang/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()
	models.SetDefaultUser()

	mux.HandleFunc("/api/v1/users", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Println("Ejecutando servidor en el puerto :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))

}
