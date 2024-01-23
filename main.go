package main

import (
	"log"
	"net/http"

	"golang-pkllaporan/config"
	authcontroller "golang-pkllaporan/controllers"
	"golang-pkllaporan/middlewares"

	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}", authcontroller.DeleteUser).Methods("DELETE")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			log.Println("Registered Route:", pathTemplate)
		}
		return nil
	})

	api := r.PathPrefix("/api").Subrouter()
	// api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
