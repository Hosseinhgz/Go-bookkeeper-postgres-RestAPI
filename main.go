package main

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	// next line is create unique id, date created and date updated and date deleted
	gorm.Model
	Name  string
	Email string `gorm:"typevarchar(100);unique_index"` // with help of gorm we make email a unique value
	Book  []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int
	PersonID   int // use name of the struct that you want make relationship

}

var db *gorm.DB
var err error

func main() {
	// LOading enviromental variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// Openning connection with database
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	// close connection to database when the main function finishes
	defer db.Close()

	// Make migrations to database if they have not already created
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})
}
