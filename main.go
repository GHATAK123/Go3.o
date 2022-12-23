package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/profile", getAllProfiles).Methods("GET")
	route.HandleFunc("/profile/{id}", getProfileById).Methods("GET")
	route.HandleFunc("/profile", addProfile).Methods("POST")
	route.HandleFunc("/profile/{id}", updateProfile).Methods("PUT")
	route.HandleFunc("/profile/{id}", deleteProfile).Methods("Delete")
	fmt.Println("Server is running at port 5000")
	http.ListenAndServe(":5000", route)

}
