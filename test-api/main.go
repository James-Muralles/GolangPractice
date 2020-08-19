package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
//Book structs (model)
 
type Book struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`

}

//Author Struct

type Author struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}




func main() {
//initialize router

r := mux.NewRouter()
r.HandleFunc("/api/books", getBooks).Methods("GET")
r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
r.HandleFunc("/api/books", createBooks).Methods("POST")
r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
r.HandleFunc("/api/books{id}", deleteBooks).Methods("DELETE")
log.Fatal(http.ListenAndServe(":8080", r))
}