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
    if err := loadEnv(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    var err error
    DB, err = setupDatabase()
    if err != nil {
        log.Fatalf("Failed to set up database: %v", err)
    }
}

func loadEnv() error {
    err := godotenv.Load()
    if err != nil {
        return fmt.Errorf("error loading .env file: %v", err)
    }
    return nil
}

func setupDatabase() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, fmt.Errorf("error opening database connection: %v", err)
    }

    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("error on database ping test: %v", err)
    }

    log.Println("Successfully connected to database!")
    return db, nil
}

func main() {
    if DB == nil {
        log.Fatal("Database connection is not initialized.")
    }

    if err := queryDatabase(); err != nil {
        log.Fatalf("Failed to query database: %v", err)
    }
}

func queryDatabase() error {
    rows, err := DB.Query("SELECT id, name FROM table_name")
    if err != nil {
        return fmt.Errorf("error querying database: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            log.Printf("Error scanning row: %v", err)
            continue
        }
        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }

    if err = rows.Err(); err != nil {
        return fmt.Errorf("error iterating over rows: %v", err)
    }

    return nil
}