package main

import (
	"github.com/gorilla/mux"
	"gorm/handlers"
	"log"
	"net/http"
)

func main() {

	// Se va a migrar cuando no exista
	// models.MigrateUser()

	// Rutas
	mux := mux.NewRouter()

	// Endpoints
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))

}
