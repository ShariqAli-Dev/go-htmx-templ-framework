package types

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int
	PersonID   int
}
