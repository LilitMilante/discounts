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
	q := `
INSERT INTO discounts (client_name, client_number, sale, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5) RETURNING client_id
`

	err := s.db.QueryRow(q, d.ClientName, d.ClientNumber, d.Sale, d.CreatedAt, d.UpdatedAt).Scan(&d.ClientID)

	return d, err
}

func (s Store) SelectClientDiscounts() ([]entity.ClientDiscount, error) {
	r, err := s.db.Query(`SELECT client_id, client_name, client_number, sale, created_at, updated_at FROM discounts`)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var ds []entity.ClientDiscount

	var d entity.ClientDiscount
	for r.Next() {
		err := r.Scan(&d.ClientID, &d.ClientName, &d.ClientNumber, &d.Sale, &d.CreatedAt, &d.UpdatedAt)
		if err != nil {
			return nil, err
		}

		ds = append(ds, d)
	}

	return ds, r.Err()
}

func (s Store) SelectClientDiscountByNumber(numb string) (entity.ClientDiscount, error) {
	var d entity.ClientDiscount

	q := `
SELECT client_id, client_name, client_number, sale, created_at, updated_at 
FROM discounts 
WHERE client_number = $1
`
	err := s.db.QueryRow(q, numb).
		Scan(&d.ClientID, &d.ClientNumber, &d.ClientNumber, &d.Sale, &d.CreatedAt, &d.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return d, domain.ErrNotFound
		}

		return d, err
	}

	return d, nil
}
