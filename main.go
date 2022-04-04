package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookKeeperRoutes(r)
	http.Handle("/", r)

	// deployed version on Heroku
	port := os.Getenv("PORT")
	log.Println("Listening on:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
	//local version
	//log.Fatal(http.ListenAndServe("localhost:9010", r))
	//fmt.Printf("App is listening on port 9010")

}
