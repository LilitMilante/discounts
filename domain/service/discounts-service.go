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

func (ds DiscountService) ClientDiscounts(f entity.ClientDiscountFilters) ([]entity.ClientDiscount, error) {
	return ds.store.SelectClientDiscounts(f)
}

func (ds DiscountService) ClientDiscountByNumber(numb string) (entity.ClientDiscount, error) {
	return ds.store.SelectClientDiscountByNumber(numb)
}

func (ds DiscountService) CreatedClientDiscount(d entity.ClientDiscount) (entity.ClientDiscount, error) {
	_, err := ds.store.SelectClientDiscountByNumber(d.ClientNumber)
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

func (ds DiscountService) EditClientDiscountByNumber(d entity.UpdateClientDiscount, number string) (entity.ClientDiscount, error) {
	_, err := ds.store.SelectClientDiscountByNumber(number)
	if err != nil {
		return entity.ClientDiscount{}, err
	}
	nowT := time.Now()
	d.UpdatedAt = &nowT

	return ds.store.UpdateClientDiscountByNumber(d, number)
}

func (ds DiscountService) DeleteClientDiscount(numb string) error {
	_, err := ds.store.SelectClientDiscountByNumber(numb)
	if err != nil {
		return err
	}

	return ds.store.DeleteClientDiscount(numb)
}
