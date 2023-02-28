package main

import "net/http"

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	app.users.All()

}

func findByID(w http.ResponseWriter, r *http.Request) {

}

func insert(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}
