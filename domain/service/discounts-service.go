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

func (ds DiscountService) CreatedClientDiscount(d entity.ClientDiscount) (entity.ClientDiscount, error) {
	_, err := ds.store.SelectClientDiscountByNumber(d.ClientNumber)

	isNotFound := errors.Is(err, domain.ErrNotFound)

	if !isNotFound && err != nil {
		return d, err
	}

	if isNotFound {
		d.CreatedAt = time.Now().UTC()
		d.UpdatedAt = time.Now().UTC()
		d.Sale = 2

		d, err = ds.store.InsertClientDiscount(d)
		if err != nil {
			return d, err
		}

		return d, nil
	}

	return d, domain.ErrDuplicateKey
}
