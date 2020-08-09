package main

import (
    "html/template"
    "net/http"
    "./model"

    _ "github.com/lib/pq"
)

// package level scope, which means it can be access anywhere in this packager.
var tpl *template.Template

func main() {
    http.HandleFunc("/", index)
http.HandleFunc("/list", model.DogIndex)
http.HandleFunc("/dog/show", model.DogShow)

http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/list", http.StatusSeeOther)
}