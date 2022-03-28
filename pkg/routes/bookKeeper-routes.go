package routes

import (
	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookKeeperRoutes = func(router *mux.Router) {
	router.HandleFunc("/persons", controllers.GetPersons).Methods("GET")
}
