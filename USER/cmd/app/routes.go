package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/", app.all).Methods("GET")
	r.HandleFunc("/api/users/{id}", app.findByID).Methods("GET")
	r.HandleFunc("/api/users/", app.insert).Methods("POST")
	r.HandleFunc("/api/users/delete", app.delete).Methods("DELETE")

}
