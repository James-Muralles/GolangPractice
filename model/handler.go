package model

import (
    "net/http"
    "../config"
    "database/sql"
)
func DogIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }
    dogs, err := ShowAllDogs()
    if err != nil {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        return
    }
    config.TPL.ExecuteTemplate(w, "index.gohtml", dogs)
}

func DogShow(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    dog, err := SingleDog(r)
    switch {
    case err == sql.ErrNoRows:
        http.NotFound(w, r)
        return
    case err != nil:
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        return
    }

    config.TPL.ExecuteTemplate(w, "show.gohtml", dog)
}