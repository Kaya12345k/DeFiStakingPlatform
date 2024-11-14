package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Successfully connected to database!")
    DB = db
}

func main() {
    rows, _ := DB.Query("SELECT id, name FROM table_name")
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        rows.Scan(&id, &name)
        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }
}