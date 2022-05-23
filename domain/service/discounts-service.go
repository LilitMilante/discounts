package service

import (
	"discounts/domain"
	"discounts/domain/entity"
	"discounts/store"
	"errors"
	"time"
)

type DiscountService struct {
	store *store.Store
}

func NewDiscountService(s *store.Store) *DiscountService {
	ds := DiscountService{
		store: s,
	}

	return &ds
}

func (ds DiscountService) Users(f entity.UserFilter) ([]entity.User, error) {
	return ds.store.SelectUsers(f)
}

func (ds DiscountService) UserByPhone(ph string) (entity.User, error) {
	return ds.store.SelectUserByPhone(ph)
}

func (ds DiscountService) CreateUser(d entity.Client) (entity.Client, error) {
	_, err := ds.store.SelectUserByPhone(d.Phone)
	switch {
	case err == nil:
		return d, domain.ErrDuplicateKey
	case !errors.Is(err, domain.ErrNotFound):
		return d, err
	}

	d.CreatedAt = time.Now()
	d.UpdatedAt = d.CreatedAt
	d.Sale = 2

	return ds.store.InsertClientDiscount(d)
}

func (ds DiscountService) EditClientDiscountByNumber(d entity.UpdateClient, number string) (entity.Client, error) {
	_, err := ds.store.SelectUserByPhone(number)
	if err != nil {
		return entity.Client{}, err
	}
	nowT := time.Now()
	d.UpdatedAt = &nowT

	return ds.store.UpdateClientDiscountByNumber(d, number)
}

func (ds DiscountService) DeleteClientDiscount(numb string) error {
	_, err := ds.store.SelectUserByPhone(numb)
	if err != nil {
		return err
	}

	return ds.store.DeleteClientDiscount(numb)
}
