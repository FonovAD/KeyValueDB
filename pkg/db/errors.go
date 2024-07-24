package db

import "errors"

var (
	ErrInvalidKey     = errors.New("error: invalid key")
	ErrInvalidValue   = errors.New("error: invalid Value")
	ErrKeyTooLong     = errors.New("error: invalid key: the key is too long(length > 100)")
	ErrRecordNotFound = errors.New("error: record not found")
)
