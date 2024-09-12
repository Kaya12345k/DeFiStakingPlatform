package main

import (
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

var dbConnStr = "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

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
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&User{}, &Product{})

	return db
}

func main() {
	db := initDB()

	newUser := User{Name: "John Doe", Email: "john.doe@example.com", Password: "supersecurepassword"}
	result := db.Create(&newUser) 
	if result.Error != nil {
		panic(result.Error)
	}
}