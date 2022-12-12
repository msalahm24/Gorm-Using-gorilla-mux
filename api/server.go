package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mahmoud24598salah/Gorm-Using-gorilla-mux.git/db"
)

func HandleReq() {
	DB := db.Init()
	handlre := New(DB)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", handlre.allUser).Methods("GET")
	router.HandleFunc("/users/{name}", handlre.deleteUser).Methods("DELETE")
	router.HandleFunc("/users/{name}/{email}",handlre.updateUser).Methods("PUT")
	router.HandleFunc("/users/{name}/{email}",handlre.newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
