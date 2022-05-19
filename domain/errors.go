package domain

import "errors"

var ErrNotFound = errors.New("client number not found")
var ErrDuplicateKey = errors.New("this client already exists")
var ErrValidate = errors.New("error validate")
