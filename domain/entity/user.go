package entity

import (
	"discounts/domain"
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

type UserRole uint8

const (
	UserRoleClient UserRole = iota
	UserRoleManager
	UserRoleDirector
	UserRoleAdmin
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      UserRole  `json:"role"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUser struct {
	Name      *string    `json:"name"`
	Phone     *string    `json:"phone"`
	Role      *UserRole  `json:"role"`
	Password  *string    `json:"password"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u User) Validate() error {
	errTxt := ""
	err := validateUserName(u.Name)
	if err != nil {
		errTxt += err.Error() + "; "
	}
	err = validateUserPhone(u.Phone)
	if err != nil {
		errTxt += err.Error()
	}
	if len(errTxt) != 0 {
		return fmt.Errorf("%w: %s", domain.ErrValidate, errTxt)
	}
	return nil
}

func (u UpdateUser) Validate() error {
	errTxt := ""

	if u.Name != nil {
		err := validateUserName(*u.Name)
		if err != nil {
			errTxt += err.Error() + "; "
		}
	}
	if u.Phone != nil {
		err := validateUserPhone(*u.Phone)
		if err != nil {
			errTxt += err.Error() + "; "
		}
	}
	if len(errTxt) != 0 {
		return fmt.Errorf("%w: %s", domain.ErrValidate, errTxt)
	}
	return nil
}

func validateUserName(name string) error {
	count := utf8.RuneCountInString(name)

	if count < 3 || count > 50 {
		return errors.New("name must be between 3 and 50 characters")
	}

	return nil
}

func validateUserPhone(num string) error {
	if !regexpNumber.MatchString(num) {
		return errors.New("wrong number entered")
	}

	return nil
}
