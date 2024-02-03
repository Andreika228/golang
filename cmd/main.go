package main

import (
	"github.com/Andreika228/awesomeProjest1/pkg/functions"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", functions.HealthCheck).Methods("GET")
	router.HandleFunc("/team/{num}", functions.GetPersonnum).Methods("GET")
	//router.HandleFunc("/team/{second_name}", functions.GetPersonName).Methods("GET")
	router.HandleFunc("/team", functions.Persons).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
