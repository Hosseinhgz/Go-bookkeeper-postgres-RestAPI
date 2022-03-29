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
	host := "ec2-99-80-170-190.eu-west-1.compute.amazonaws.com"
	dbPort := "5432"
	user := "arfijkzivxdicj"
	dbName := "dfki72t99bpp2i"
	password := "f2a30019890da78a4fd8fc8a121878fb8685084a27ad0b0337d4a90c698e822f"

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
