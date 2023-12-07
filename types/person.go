package types

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100)"`
	Books []Book
}
