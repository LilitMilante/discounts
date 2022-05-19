package entity

import (
	"discounts/domain"
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"
)

var regexpNumber = regexp.MustCompile(`^(\+375)(25|29|33|44)\d{7}$`)

type ClientDiscount struct {
	ClientID     int64     `json:"client_id"`
	ClientName   string    `json:"client_name"`
	ClientNumber string    `json:"client_number"`
	Sale         int8      `json:"sale"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateClientDiscount struct {
	ClientName   *string    `json:"client_name"`
	ClientNumber *string    `json:"client_number"`
	Sale         *int8      `json:"sale"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func (cd ClientDiscount) Validate() error {
	err := validClientName(cd.ClientName)
	if err != nil {
		return err
	}
	err = validClientNumber(cd.ClientNumber)
	if err != nil {
		return err
	}
	return nil
}

func (ucd UpdateClientDiscount) Validate() error {
	if ucd.ClientName != nil {
		err := validClientName(*ucd.ClientName)
		if err != nil {
			return err
		}
	}
	if ucd.ClientNumber != nil {
		err := validClientNumber(*ucd.ClientNumber)
		if err != nil {
			return err
		}
	}
	if ucd.Sale != nil {
		err := validClientSale(*ucd.Sale)
		if err != nil {
			return err
		}
	}
	return nil
}

func validClientName(name string) error {
	count := utf8.RuneCountInString(name)

	if count < 3 || count > 50 {
		return fmt.Errorf("%w: name must be between 3 and 50 characters", domain.ErrValidate)
	}

	return nil
}

func validClientNumber(num string) error {
	if !regexpNumber.MatchString(num) {
		return fmt.Errorf("%w: wrong number entered", domain.ErrValidate)
	}

	return nil
}

func validClientSale(sale int8) error {
	if sale < 0 || sale > 50 {
		errTxt := "discount must be at least 0% and not more than 50%"
		return fmt.Errorf("%w: %s", domain.ErrValidate, errTxt)
	}

	return nil
}
