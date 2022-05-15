package store

import (
	"database/sql"
	"discounts/domain"
	"discounts/domain/entity"
	"errors"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s Store) InsertClientDiscount(d entity.ClientDiscount) (entity.ClientDiscount, error) {
	err := s.db.QueryRow(
		`INSERT INTO discounts (client_name, client_number, sale, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5) RETURNING client_id`,
		d.ClientName, d.ClientNumber, d.Sale, d.CreatedAt, d.UpdatedAt,
	).Scan(&d.ClientID)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (s Store) SelectClientDiscountByNumber(numb string) (entity.ClientDiscount, error) {
	var d entity.ClientDiscount

	err := s.db.QueryRow(`SELECT client_id, client_name, client_number, sale, created_at, updated_at 
FROM discounts
WHERE client_number = $1`, numb).Scan(&d.ClientID, &d.ClientNumber, &d.ClientNumber, &d.Sale, &d.CreatedAt, &d.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return d, domain.ErrNotFound
	}

	if err != nil {
		return d, err
	}

	return d, nil
}
