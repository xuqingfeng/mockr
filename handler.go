package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func UserIndexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "home")
}

func UserRHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuid := vars["uuid"]

}

func UserResponseHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

	case "POST":

	case "PUT":
	// check authorization

	case "DELETE":
	// check authorization
	default:

	}
}
