package domain

import "errors"

var ErrNotFound = errors.New("id not found")
var ErrDuplicateKey = errors.New("this client already exists")
