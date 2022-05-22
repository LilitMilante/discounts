package entity

import (
	"discounts/domain"
	"errors"
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

type ClientDiscountFilters struct {
	Name  *string
	Sale  *int8
	Start *time.Time
	End   *time.Time
}

func (cd ClientDiscount) Validate() error {
	errTxt := ""
	err := validClientName(cd.ClientName)
	if err != nil {
		errTxt += err.Error() + "; "
	}
	err = validClientNumber(cd.ClientNumber)
	if err != nil {
		errTxt += err.Error()
	}
	if len(errTxt) != 0 {
		return fmt.Errorf("%w: %s", domain.ErrValidate, errTxt)
	}
	return nil
}

func (ucd UpdateClientDiscount) Validate() error {
	errTxt := ""

	if ucd.ClientName != nil {
		err := validClientName(*ucd.ClientName)
		if err != nil {
			errTxt += err.Error() + "; "
		}
	}
	if ucd.ClientNumber != nil {
		err := validClientNumber(*ucd.ClientNumber)
		if err != nil {
			errTxt += err.Error() + "; "
		}
	}
	if ucd.Sale != nil {
		err := validClientSale(*ucd.Sale)
		if err != nil {
			errTxt += err.Error() + "; "
		}
	}
	if len(errTxt) != 0 {
		return fmt.Errorf("%w: %s", domain.ErrValidate, errTxt)
	}
	return nil
}

func validClientName(name string) error {
	count := utf8.RuneCountInString(name)

	if count < 3 || count > 50 {
		return errors.New("name must be between 3 and 50 characters")
	}

	return nil
}

func validClientNumber(num string) error {
	if !regexpNumber.MatchString(num) {
		return errors.New("wrong number entered")
	}

	return nil
}

func validClientSale(sale int8) error {
	if sale < 0 || sale > 50 {
		return errors.New("discount must be at least 0% and not more than 50%")
	}

	return nil
}
