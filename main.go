package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/models"
	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/routes"

)




var (
	persons = []models.Person{
		{Name: "Hossein", Email: "hossein@gmail.com"},
		{Name: "Lia", Email: "lia1989@gmail.com"},
	}
	books = []models.Book{
		{Title: "Lord of Glory", Author: "Benjamin B. Warfield", CallNumber: 238268289, PersonID: 1},
		{Title: "The Lord of Death and the Queen of Life", Author: "Homer Eon Flint", CallNumber: 234568288, PersonID: 1},
		{Title: "The Best Horror of the Year: Volume 1", Author: "Ellen Datlow", CallNumber: 762365490, PersonID: 2},
		{Title: "Finnish the Language", Author: "Unknown", CallNumber: 1236176, PersonID: 2},
	}
)

var db *gorm.DB
var err error

func main() {
	// // Loading enviromental variables - .env file is not working
	// dialect := os.Getenv("DIALECT")
	// host := os.Getenv("HOST")
	// dbPort := os.Getenv("DBPORT")
	// user := os.Getenv("USER")
	// dbName := os.Getenv("NAME")
	// password := os.Getenv("PASSWORD")
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

	// close connection to database when the main function finishes
	defer db.Close()

	// Make migrations to database if they have not already created
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Book{})

	// // creating some dummy data for testing connection with db
	// for idx := range persons {
	// 	db.Create(&persons[idx])
	// }

	// for idx := range books {
	// 	db.Create(&books[idx])
	// }


	r := mux.NewRouter()
	routes.RegisterBookKeeperRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
