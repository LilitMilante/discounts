package store

import (
	"database/sql"
	"discounts/domain"
	"discounts/domain/entity"
	"errors"
	"fmt"
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

	q := `SELECT client_id, client_name, client_number, sale, created_at, updated_at FROM discounts WHERE client_number = $1`
	err := s.db.QueryRow(q, numb).
		Scan(&d.ClientID, &d.ClientName, &d.ClientNumber, &d.Sale, &d.CreatedAt, &d.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return d, domain.ErrNotFound
		}

		return d, err
	}

	return d, nil
}

func (s Store) UpdateClientDiscountByNumber(d entity.UpdateClientDiscount, numb string) (entity.ClientDiscount, error) {
	var dc entity.ClientDiscount

	var columns string
	var args []interface{}

	if d.ClientName != nil {
		args = append(args, d.ClientName)
		columns += fmt.Sprintf("client_name = $%d, ", len(args))
	}
	if d.ClientNumber != nil {
		args = append(args, d.ClientNumber)
		columns += fmt.Sprintf("client_number = $%d, ", len(args))
	}
	if d.Sale != nil {
		args = append(args, d.Sale)
		columns += fmt.Sprintf("sale = $%d, ", len(args))
	}
	if d.UpdatedAt != nil {
		args = append(args, d.UpdatedAt)
		columns += fmt.Sprintf("updated_at = $%d, ", len(args))
	}

	args = append(args, numb)

	set := columns[:len(columns)-2] // убираем в конце запятую

	q := fmt.Sprintf("UPDATE discounts SET %s WHERE client_number = $%d RETURNING client_id, client_name, client_number, sale, created_at, updated_at", set, len(args))

	err := s.db.QueryRow(q, args...).
		Scan(&dc.ClientID, &dc.ClientName, &dc.ClientNumber, &dc.Sale, &dc.CreatedAt, &dc.UpdatedAt)
	if err != nil {
		return entity.ClientDiscount{}, err
	}

	return dc, nil
}

func (s Store) DeleteClientDiscount(numb string) error {
	_, err := s.db.Exec("DELETE FROM discounts WHERE client_number = $1", numb)

	return err
}
