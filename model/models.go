package model


import "../config"

// Fields first letter must be uppercase so we can export them to templates.
type Dog struct {
    ID         int
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