package config

import (
	"fmt"
	"log"

	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	// local version
	// dialect := "postgres"
	// host := "localhost"
	// dbPort := "5432"
	// user := "postgres"
	// dbName := "book_keeper"
	// password := ""

	// deployed version Heroku
	dialect := "postgres"
	host := "ec2-52-30-67-143.eu-west-1.compute.amazonaws.com"
	dbPort := "5432"
	user := "fepuzbklejrydh"
	dbName := "d1a6j2p9b55cbk"
	password := "2b5141aad07ec369afa35e1e173ab964529665dab2b4f31afb2e10a2ea593c5f"

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, user, dbName, password, dbPort)

	// Openning connection with database
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}
	// Make migrations to database if they have not already created
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Book{})
	return db
}
