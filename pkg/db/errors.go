package db

import "errors"

var ErrInvalidKey = errors.New("invalid key")
var ErrInvalidValue = errors.New("invalid Value")
var ErrKeyTooLong = errors.New("invalid key: the key is too long(length > 100)")
