package model


import (
    "../config"
    "net/http"
    "errors"
)

// Fields first letter must be uppercase so we can export them to templates.
type Dog struct {
    ID         string
    Name       string
    Breed	   string
}

 //Show all dogs
func ShowAllDogs() ([]Dog, error) {

    rows, err := config.DB.Query("SELECT * FROM dogs")

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    dogs := make([]Dog, 0)
    for rows.Next() {
        dog := Dog{}
        err := rows.Scan(&dog.ID, &dog.Name, &dog.Breed)
        if err != nil {

            return nil, err
        }
        dogs = append(dogs, dog)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return dogs, nil
}

func SingleDog(r *http.Request) (Dog, error) {
    dog := Dog{}
    ID := r.FormValue("id")
    if ID == "" {
        return dog, errors.New("400. Bad Request.")
    }

    row := config.DB.QueryRow("SELECT * FROM dogs WHERE id = $1", ID)

    err := row.Scan(&dog.ID, &dog.Name, &dog.Breed)
    if err != nil {
        return dog, err
    }
    return dog, nil
}

func InsertDog(r *http.Request) (Dog, error) {
    // get form values
    d := Dog{}
    d.ID = r.FormValue("dogID")
    d.Name = r.FormValue("name")
    d.Breed = r.FormValue("breed")

    
    // validate form values
    if d.ID == "" || d.Name == "" || d.Breed == ""  {
        return d, errors.New("400. Bad request. All fields must be complete.")
    }

    // insert values
    _, err := config.DB.Exec("INSERT INTO dogs (id, name, breed) VALUES ($1, $2, $3)", d.ID, d.Name, d.Breed)
    if err != nil {
        return d, errors.New("500. Internal Server Error." + err.Error())
    }
    return d, nil
}

func UpdateDog(r *http.Request) (Dog, error) {
    // get form values
    d := Dog{}
    d.ID = r.FormValue("dogID")
    d.Name = r.FormValue("name")
    d.Breed = r.FormValue("breed")
    // validate form values
    if d.ID == "" || d.Name == "" || d.Breed == "" {
        return d, errors.New("400. Bad request. All fields must be complete.")
    }

    // insert values
    _, err := config.DB.Exec("UPDATE dogs SET id= $1, name=$2, breed=$3 WHERE id=$1;", d.ID, d.Name, d.Breed)
    if err != nil {
        return d, err
    }
    return d, nil
}

func DeleteDog(r *http.Request) error {
    pID := r.FormValue("id")
    if pID == "" {
        return errors.New("400. Bad Request.")
    }

    _, err := config.DB.Exec("DELETE FROM dogs WHERE id=$1;", pID)
    if err != nil {
        return errors.New("500. Internal Server Error")
    }
    return nil
}