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
        log.Fatalf("Error loading .env file: %v", err) // Improved logging to include the error
    }

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err) // Improved logging
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error on database ping test: %v", err) // Improved logging
    }

    log.Println("Successfully connected to database!")
    DB = db
}

func main() {
    rows, err := DB.Query("SELECT id, name FROM table_name")
    if err != nil {
        log.Fatalf("Error querying database: %v", err) // Error handling for query operation
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            log.Printf("Error scanning row: %v", err) // Logging an error instead of fatal, to proceed with next rows
            continue
        }
        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }

    // Check for errors from iterating over rows
    if err = rows.Err(); err != nil {
        log.Fatalf("Error iterating over rows: %v", err)
    }
}