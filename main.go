package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	// user router :8000
	userRouter := mux.NewRouter()

	// root path
	userRootSubRouter := userRouter.PathPrefix("/").Subrouter()
	userRootSubRouter.HandleFunc("/", UserIndexHandler)
	userRootSubRouter.HandleFunc("/r/"+`{uuid:\S+}`, UserRHandler)

	// api path
	userApiSubRouter := userRouter.PathPrefix("/api").Subrouter()
	userApiSubRouter.HandleFunc("/response", UserResponseHandler)

	go http.ListenAndServe(":8000", userRouter)

	// admin router :8001
	adminRouter := mux.NewRouter()

	go http.ListenAndServe(":8001", adminRouter)
}
