package db

import (
	"sync"

	linkedlist "github.com/PepsiKingIV/KeyValueDB/pkg/db/linked_list"
)

type Store interface {
	Put(string, string) error
	Get(string) (string, error)
	Delete(string) error
	Hash(string) (int, error)
	Lock() error
	Unlock() error
	RLock() error
	RUnlock() error
}

type DB struct {
	mu     sync.RWMutex
	dbSize int
	// заменить массив из массивов на массив из указателей на связный список
	arrayOfPointers []*linkedlist.Node
}

func NewDB() *DB {
	db := DB{
		mu:              sync.RWMutex{},
		dbSize:          100,
		arrayOfPointers: make([]*linkedlist.Node, 100),
	}
	for i := range db.arrayOfPointers {
		db.arrayOfPointers[i] = linkedlist.NewLinkedList()
	}
	return &db
}

func (db *DB) Hash(key string) (int, error) {
	switch {
	case (len(key) < 1):
		return -1, ErrInvalidKey
	case (len(key) > 100):
		return -1, ErrKeyTooLong
	}
	var sum int = 0
	for _, j := range []byte(key) {
		sum += int(j)
	}
	return sum % db.dbSize, nil
}

func (db *DB) Lock() error {
	db.mu.Lock()
	return nil
}

func (db *DB) Unlock() error {
	db.mu.Unlock()
	return nil
}

func (db *DB) RLock() error {
	db.mu.RLock()
	return nil
}

func (db *DB) RUnlock() error {
	db.mu.RUnlock()
	return nil
}

func (db *DB) Put(key string, value string) error {
	switch {
	case (len(key) < 1):
		return ErrInvalidKey
	case (len(value) < 1):
		return ErrInvalidValue
	}
	ind, err := db.Hash(key)
	if err != nil {
		return ErrInvalidKey
	}
	linkedlist.Add(db.arrayOfPointers[ind], key, value)
	return nil
}

// TODO: добавить передачу по ссылке переменной для записи value
func (db *DB) Get(key string) (string, error) {
	switch {
	case (len(key) < 1):
		return "", ErrInvalidKey
	}
	ind, err := db.Hash(key)
	if err != nil {
		return "", ErrInvalidKey
	}
	rec, err := linkedlist.Get(db.arrayOfPointers[ind], key)
	if err != nil {
		return "", err
	}
	return rec.Value, err
}

func (db *DB) Delete(key string) error {
	switch {
	case (len(key) < 1):
		return ErrInvalidKey
	}
	ind, err := db.Hash(key)
	if err != nil {
		return err
	}
	err = linkedlist.Delete(db.arrayOfPointers[ind], key)
	switch {
	case (err == linkedlist.ErrNotFound):
		return ErrRecordNotFound
	case (err != nil):
		return err
	}
	return nil
}
