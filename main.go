package main

import (
	"log"
	"net/http"

	"golang-pkllaporan/config"
	"golang-pkllaporan/middlewares"

	"golang-pkllaporan/controllers/authcontroller"
	"golang-pkllaporan/controllers/productcontroller"

	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
