package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type GormStore struct {
	client *gorm.DB
}

type Store struct {
	PersonStore GormStore
}

var DB_URL string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env files")
	}
	DB_URL = os.Getenv("DB_URL")
}

// func HandleGetPeople(store *gorm.DB) (*[]Person, error) {
// 	var people []types.Person
// 	if err := store.Find(&people).Error; err != nil {
// 		return nil, err
// 	}
// 	return &people, nil
// }
