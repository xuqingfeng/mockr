package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	finish := make(chan bool)

	// user router :8000
	userRouter := mux.NewRouter()

	// web path
	userWebSubRouter := userRouter.PathPrefix("/").Subrouter()
	userWebSubRouter.HandleFunc("/", UserIndexHandler)
	userWebSubRouter.HandleFunc("/r", UserRHandler)
	userWebSubRouter.HandleFunc("/r/"+`{uuid:\S+}`, UserRUUIDHandler)

	// api path
	userApiSubRouter := userRouter.PathPrefix("/api").Subrouter()
	userApiSubRouter.HandleFunc("/response/"+`{uuid:\S+}`, UserResponseHandler)

	// assets path
	userRouter.HandleFunc("/assets/"+`{path:\S+}`, AssetsHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8000", userRouter))
	}()

	// admin router :8001
	adminRouter := mux.NewRouter()

	go func() {
		log.Fatal(http.ListenAndServe(":8001", adminRouter))
	}()

	<-finish
}
