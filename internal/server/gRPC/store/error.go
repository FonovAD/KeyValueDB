package store

import "errors"

var (
	ErrStorePut = errors.New("store: data recording error")
	ErrStoreGet = errors.New("store: data reading error")
	ErrStoreDel = errors.New("store: data deletion error")
)
