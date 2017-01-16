package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	finish := make(chan bool)

	// user router :8000
	userRouter := mux.NewRouter()

	// web path
	userWebSubRouter := userRouter.PathPrefix("/").Subrouter()
	userWebSubRouter.HandleFunc("/", UserIndexHandler)
	userWebSubRouter.HandleFunc("/r/"+`{uuid:\S+}`, UserRHandler)

	// api path
	userApiSubRouter := userRouter.PathPrefix("/api").Subrouter()
	userApiSubRouter.HandleFunc("/response", UserResponseHandler)

	// assets path
	userRouter.HandleFunc("/assets/"+`{path:\S+}`, AssetsHandler)

	go func() {
		http.ListenAndServe(":8000", userRouter)
	}()

	// admin router :8001
	adminRouter := mux.NewRouter()

	go func() {
		http.ListenAndServe(":8001", adminRouter)
	}()

	<-finish
}
