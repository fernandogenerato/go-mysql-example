package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-mysql-example/controller"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.FindUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8899", router))

}
