package main

import (
	"github.com/jinzhu/gorm"
)

type Person struct{
	// next line is create unique id, date created and date updated and date deleted
	gorm.Model
	Name string
	Email string `gorm:"typevarchar(100);unique_index"` // with help of gorm we make email a unique value
	Book []Book
}

type Book struct {
	gorm.Model

	Title string
	Author string
	CallNumber int
	PersonID int // use name of the struct that you want make relationship

}

func main(){
	
}