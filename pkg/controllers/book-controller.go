package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/models"
	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/utils"
	"github.com/gorilla/mux"
)

// router structure:
// router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
// router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
// router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBooks()
	// res is json version of the result that we get from database
	res, _ := json.Marshal(NewBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// "bookId" is come from path param
	bookId := vars["bookId"]

	// bookId in url is string here we change it to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing - GetBookById()")
	}

	// because models.GetBookById is returns 2 item and we dont want 2nd item in this func we use _
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}

	// we getting some data in json from user in the request
	// here we should change it to sth that db can understand
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing - DeleteBook()")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var editedBook = &models.Book{}
	// editedBook is created from request body which contains json file with required values
	utils.ParseBody(r, editedBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing - UpdateBook()")
	}
	// selectedBookdb is book inside database that we want to update
	selectedBookdb, db := models.GetBookById(Id)
	if editedBook.Name != "" {
		selectedBookdb.Name = editedBook.Name
	}
	if editedBook.Author != "" {
		selectedBookdb.Author = editedBook.Author
	}
	if editedBook.Publication != "" {
		selectedBookdb.Publication = editedBook.Publication
	}
	db.Save(&selectedBookdb)
	res, _ := json.Marshal(selectedBookdb)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
