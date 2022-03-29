package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/config"
	"github.com/Hosseinhgz/Go-bookkeeper-postgres-RestAPI/pkg/models"
	"github.com/gorilla/mux"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	var persons []models.Person
	db := config.GetDB()
	db.Find(&persons)
	json.NewEncoder(w).Encode(&persons)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	db := config.GetDB()
	db.Find(&books)
	json.NewEncoder(w).Encode(&books)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := config.GetDB()

	var person models.Person
	var books []models.Book
	db.First(&person, params["id"])
	db.Model(&person).Related(&books) // get all the books which is related to that person

	person.Books = books
	json.NewEncoder(w).Encode(person)
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := config.GetDB()

	var book models.Book
	db.First(&book, params["id"])
	json.NewEncoder(w).Encode(book)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	db := config.GetDB()
	json.NewDecoder(r.Body).Decode(&person)
	createdPerson := db.Create(&person)
	err := createdPerson.Error
	if err !=nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(createdPerson)
	}
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	db := config.GetDB()
	json.NewDecoder(r.Body).Decode(&book)
	createdBook := db.Create(&book)
	err := createdBook.Error
	if err !=nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(createdBook)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := config.GetDB()

	var person models.Person
	db.Delete(&person, params["id"])
	json.NewEncoder(w).Encode(person)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := config.GetDB()

	var book models.Book
	db.Delete(&book, params["id"])
	json.NewEncoder(w).Encode(book)
}

//*********FROM PREVIOUS APP**************************************************START
// var NewBook models.Book

// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	NewBooks := models.GetAllBooks()
// 	// res is json version of the result that we get from database
// 	res, _ := json.Marshal(NewBooks)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetBookById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	// "bookId" is come from path param
// 	bookId := vars["bookId"]

// 	// bookId in url is string here we change it to int
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing - GetBookById()")
// 	}

// 	// because models.GetBookById is returns 2 item and we dont want 2nd item in this func we use _
// 	bookDetails, _ := models.GetBookById(ID)
// 	res, _ := json.Marshal(bookDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	createBook := &models.Book{}

// 	// we getting some data in json from user in the request
// 	// here we should change it to sth that db can understand
// 	utils.ParseBody(r, createBook)
// 	b := createBook.CreateBook()
// 	res, _ := json.Marshal(b)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing - DeleteBook()")
// 	}
// 	book := models.DeleteBook(ID)
// 	res, _ := json.Marshal(book)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	var editedBook = &models.Book{}
// 	// editedBook is created from request body which contains json file with required values
// 	utils.ParseBody(r, editedBook)

// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	Id, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing - UpdateBook()")
// 	}
// 	// selectedBookdb is book inside database that we want to update
// 	selectedBookdb, db := models.GetBookById(Id)
// 	if editedBook.Name != "" {
// 		selectedBookdb.Name = editedBook.Name
// 	}
// 	if editedBook.Author != "" {
// 		selectedBookdb.Author = editedBook.Author
// 	}
// 	if editedBook.Publication != "" {
// 		selectedBookdb.Publication = editedBook.Publication
// 	}
// 	db.Save(&selectedBookdb)
// 	res, _ := json.Marshal(selectedBookdb)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
//*********FROM PREVIOUS APP**************************************************END
