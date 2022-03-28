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
	dialect := "postgres"
	host := "localhost"
	dbPort := "5432"
	user := "postgres"
	dbName := "book_keeper"
	password := "Hhh44974497"

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

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
