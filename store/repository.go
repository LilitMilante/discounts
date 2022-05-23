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

func (s Store) InsertClientDiscount(d entity.Client) (entity.Client, error) {
	q := `
INSERT INTO discounts (client_name, client_number, sale, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5) RETURNING client_id
`

	err := s.db.QueryRow(q, d.Name, d.Phone, d.Sale, d.CreatedAt, d.UpdatedAt).Scan(&d.ID)

	return d, err
}

func (s Store) SelectUsers(f entity.UserFilter) ([]entity.User, error) {
	var columns string
	var args []interface{}

	if f.Name != nil {
		args = append(args, f.Name)
		columns += fmt.Sprintf("name = $%d AND ", len(args))
	}
	if f.Sale != nil {
		args = append(args, f.Sale)
		columns += fmt.Sprintf("sale = $%d AND ", len(args))
	}
	if f.Start != nil && f.End != nil {
		args = append(args, f.Start, f.End)
		columns += fmt.Sprintf("created_at BETWEEN $%d AND $%d AND ", len(args)-1, len(args))
	}

	q := "SELECT id, name, phone, role, login, created_at, updated_at FROM users "

	if len(args) != 0 {
		where := columns[:len(columns)-5] // убираем в конце пробелы и AND
		q += fmt.Sprintf(" WHERE %s", where)
	}

	r, err := s.db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var us []entity.User

	var u entity.User
	for r.Next() {
		err := r.Scan(&u.ID, &u.Name, &u.Phone, &u.Role, &u.Login, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		us = append(us, u)
	}

	return us, r.Err()
}

func (s Store) SelectUserByPhone(ph string) (entity.User, error) {
	var u entity.User

	q := `SELECT id, name, phone, role, login, created_at, updated_at FROM users WHERE phone = $1`
	err := s.db.QueryRow(q, ph).
		Scan(&u.ID, &u.Name, &u.Phone, &u.Role, &u.Login, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, domain.ErrNotFound
		}

		return u, err
	}

	return u, nil
}

func (s Store) UpdateClientDiscountByNumber(d entity.UpdateClient, numb string) (entity.Client, error) {
	var dc entity.Client

	var columns string
	var args []interface{}

	if d.Name != nil {
		args = append(args, d.Name)
		columns += fmt.Sprintf("client_name = $%d, ", len(args))
	}
	if d.Phone != nil {
		args = append(args, d.Phone)
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

	set := columns[:len(columns)-2] // убираем в конце пробел и запятую

	q := fmt.Sprintf("UPDATE discounts SET %s WHERE client_number = $%d RETURNING client_id, client_name, client_number, sale, created_at, updated_at", set, len(args))

	err := s.db.QueryRow(q, args...).
		Scan(&dc.ID, &dc.Name, &dc.Phone, &dc.Sale, &dc.CreatedAt, &dc.UpdatedAt)
	if err != nil {
		return entity.Client{}, err
	}

	return dc, nil
}

func (s Store) DeleteClientDiscount(numb string) error {
	_, err := s.db.Exec("DELETE FROM discounts WHERE client_number = $1", numb)

	return err
}
