package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
	"strings"
)

var (
	mimeTypes = map[string]string{
		".css": "text/css",
		".js":  "text/javascript",
	}
)

// user handler
func UserIndexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	asset, err := Asset("assets/web/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, string(asset))
}

func UserRHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuid := vars["uuid"]
	fmt.Fprint(w, uuid)
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

// admin handler

// utility
func AssetsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	asset, err := Asset("assets/" + vars["path"])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	ctype := mime.TypeByExtension(vars["path"])
	if ctype == "" {
		// use extra mimeTypes
		for key, value := range mimeTypes {
			if strings.HasSuffix(vars["path"], key) {
				ctype = value
				break
			}
		}
	}
	w.Header().Set("Content-Type", ctype)
	fmt.Fprint(w, string(asset))
}
