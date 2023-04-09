package store

import (
	e "github.com/David-Orson/nb/error"
	"gorm.io/gorm"
)

type Actions interface {
	Account() AccountActions
}

type Store struct {
	DB *gorm.DB
}

func New() (s *Store) {
	s = &Store{}

	if !e.Process(s.newDB(), "Error creating database connection: ") {
		return
	}

	e.Process(s.migrations(), "Error running migrations: ")

	return
}

func (s *Store) Account() AccountActions {
	return &accountStore{s}
}
