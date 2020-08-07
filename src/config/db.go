var DB *sql.DB

func init() {
   var err error
    DB, err = sql.Open("postgres", "postgres://postgres:codeup@localhost/test_db?sslmode=disable")
    if err != nil {
        panic(err)
    }

    if err = DB.Ping(); err != nil {
        panic(err)
    }
    fmt.Println("Database connection is succesful")

}