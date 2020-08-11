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

    config.TPL.ExecuteTemplate(w, "dog.gohtml", dog)
}

func DogCreate(w http.ResponseWriter, r *http.Request) {
    config.TPL.ExecuteTemplate(w, "createForm.gohtml", nil)
}

func DogCreateProcess(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    dog, err := InsertDog(r)
    if err != nil {
        http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
        return
    }

    config.TPL.ExecuteTemplate(w, "dog.gohtml", dog)
}

func DogUpdate(w http.ResponseWriter, r *http.Request) {
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

    config.TPL.ExecuteTemplate(w, "updateDog.gohtml", dog)
}

func DogUpdateProcess(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    dog, err := UpdateDog(r)
    if err != nil {
        http.Error(w, http.StatusText(406), http.StatusBadRequest)
        return
    }

    config.TPL.ExecuteTemplate(w, "updatedDog.gohtml", dog)
}