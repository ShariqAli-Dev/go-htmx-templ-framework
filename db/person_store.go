package db

import (
	"github.com/jinzhu/gorm"
	"github.com/shariqali-dev/quizmify/types"
)

type PersonStore interface {
	GetPeople() (*[]types.Person, error)
}

func NewPersonStore(client *gorm.DB) *GormStore {
	return &GormStore{
		client: client,
	}
}

func (s *GormStore) GetPeople() (*[]types.Person, error) {
	var people []types.Person
	if err := s.client.Find(&people).Error; err != nil {
		return nil, err
	}
	return &people, nil
}
