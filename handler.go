package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xuqingfeng/mockr/response"
	"github.com/xuqingfeng/mockr/util"
	"mime"
	"net/http"
	"strconv"
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
		return
	}
	fmt.Fprint(w, string(asset))
}

func UserRHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	asset, err := Asset("assets/web/r.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprint(w, string(asset))
}

func UserRUUIDHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	fmt.Fprint(w, uuid)
}

func UserResponseHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

	case "POST":

		var rp response.Response
		err := json.NewDecoder(r.Body).Decode(&rp)
		if err != nil {
			sendMessage(w, false, "E! "+err.Error(), new(struct{}))
		} else {
			err = rp.AddResponse()
			if err != nil {
				sendMessage(w, false, "E! "+err.Error(), new(struct{}))
			} else {
				sendMessage(w, true, "I! add response success", rp)
			}
		}
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
		return
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
	// cache - Last-Modified, ETag
	w.Header().Set("Cache-Control", "max-age="+strconv.Itoa(24*60*60))

	fmt.Fprint(w, string(asset))
}

func sendMessage(w http.ResponseWriter, success bool, message string, data interface{}) {

	msg := util.Msg{
		Success: success,
		Message: message,
		Data:    data,
	}

	msgInByteSlice, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msgInByteSlice)
}
