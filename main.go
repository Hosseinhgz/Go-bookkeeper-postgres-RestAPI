package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/models"
	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/routes"
	"github.com/gorilla/mux"
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

func main() {
	r := mux.NewRouter()
	routes.RegisterBookKeeperRoutes(r)
	http.Handle("/", r)

	// //local version
	// log.Fatal(http.ListenAndServe("localhost:9010", r))

	// deployed version on Heroku
	port := ":" + os.Getenv("PORT")
	log.Println("Listening on:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
