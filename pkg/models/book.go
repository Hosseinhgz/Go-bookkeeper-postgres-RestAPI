package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int
	PersonID   int // use name of the struct that you want make relationship

}
