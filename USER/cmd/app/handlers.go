package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	users, err := app.users.All()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(users)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("content-tyep", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.users.FindbyID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("content-type", "application")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
