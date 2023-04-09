package store

import (
	"gorm.io/gorm"
	e "nautilus-billing.com/m/v2/error"
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
