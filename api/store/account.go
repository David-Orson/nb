package store

import (
	e "github.com/David-Orson/nb/error"
	"github.com/David-Orson/nb/store/model"
)

type AccountActions interface {
	Create(acc *model.Account) (ok bool)
}

type accountStore struct {
	*Store
}

func (s *accountStore) Create(acc *model.Account) (ok bool) {
	return e.Process(s.DB.Create(acc).Error, "Error creating account: ")
}
