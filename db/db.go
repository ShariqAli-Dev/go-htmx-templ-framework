package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100)"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int
	PersonID   int
}

var person = &Person{
	Name:  "shairq",
	Email: "email@email.com",
	Books: []Book{
		{Title: "the coolest book ever", Author: "teste", CallNumber: 234, PersonID: 10},
	},
}

func main() {
	client, err := gorm.Open("postgres", "postgres://postgres:password@localhost:5432/quizmify")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// database migrations
	client.AutoMigrate(&Person{})
	client.AutoMigrate(&Book{})

	people, err := HandleGetPeople(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("this is the people\n", people)

}

func HandleGetPeople(store *gorm.DB) (*[]Person, error) {
	var people []Person
	if err := store.Find(&people).Error; err != nil {
		return nil, err
	}
	return &people, nil
}
