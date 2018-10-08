package main

import (
	"go-microservice/handle"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", handle.UserProfile).Queries("access_token", "{access_token}").Methods("GET")
	router.HandleFunc("/login", handle.Login).Methods("POST")
	router.HandleFunc("/register", handle.RegisterNewUser).Methods("POST")
	// router.HandleFunc("/", handle.NOTIMPLENTED).Methods("GET")
	// router.HandleFunc("/HOLD/{hold}", handle.NOTIMPLENTED).Methods("PUT")
	http.ListenAndServe(":8081", router)
}
