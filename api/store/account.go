package store

import (
	e "nautilus-billing.com/m/v2/error"
	"nautilus-billing.com/m/v2/store/model"
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
