package main

import (
    "fmt"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    dbHost     = os.Getenv("DB_HOST")
    dbPort     = os.Getenv("DB_PORT")
    dbUser     = os.Getenv("DB_USER")
    dbPassword = os.Getenv("DB_PASSWORD")
    dbName     = os.Getenv("DB_NAME")
)

var dbConnStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    dbHost, dbPort, dbUser, dbPassword, dbName)

type User struct {
    gorm.Model
    Name     string
    Email    string `gorm:"unique"`
    Password string
}

type Product struct {
    gorm.Model
    Name        string
    Description string
    Price       float64
}

func initDB() *gorm.DB {
    db, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to database: %v", err))
    }
    if err := db.AutoMigrate(&User{}, &Product{}); err != nil {
        panic(fmt.Sprintf("Failed to automatically migrate database schema: %v", err))
    }

    return db
}

func main() {
    db := initDB()

    newUser := User{Name: "John Doe", Email: "john.doe@example.com", Password: "supersecurepassword"}
    result := db.Create(&newUser)
    if result.Error != nil {
        panic(fmt.Sprintf("Failed to create new user: %v", result.Error))
    }
}