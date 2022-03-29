package routes

import (
	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookKeeperRoutes = func(router *mux.Router) {
	router.HandleFunc("/persons", controllers.GetPersons).Methods("GET")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.GetPerson).Methods("GET")              // get specific person and all his books
	router.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")              // get specific book
	router.HandleFunc("/create/person", controllers.CreatePerson).Methods("POST")        // Create new person in db
	router.HandleFunc("/delete/person/{id}", controllers.DeletePerson).Methods("DELETE")
}
