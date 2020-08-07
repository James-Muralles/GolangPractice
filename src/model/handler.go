package model

import (
    "net/http"
    "src/config"
)
func DogIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }
    patients, err := ShowAllDogs()
    if err != nil {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
        return
    }
    config.TPL.ExecuteTemplate(w, "index.gohtml", dogs)
}