package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/image", imageRegistration).Methods("POST")
	router.HandleFunc("/image/{sha}/chunks", uploadImage).Methods("POST")
	router.HandleFunc("/image/{sha}", downloadImage).Methods("GET")
	log.Fatal(http.ListenAndServe(":4444", router))
}