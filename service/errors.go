package service

import (
	"errors"
)

var (
	ErrDuplicateEmail = errors.New("user already exists")
)
