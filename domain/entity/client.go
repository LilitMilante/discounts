package entity

import (
	"regexp"
	"time"
)

var regexpNumber = regexp.MustCompile(`^(\+375)(25|29|33|44)\d{7}$`)

type Client struct {
	User
	Sale int8 `json:"sale"`
}

type UpdateClient struct {
	UpdateUser
	Sale *int8 `json:"sale"`
}

type UserFilter struct {
	Name  *string
	Sale  *int8
	Start *time.Time
	End   *time.Time
}
