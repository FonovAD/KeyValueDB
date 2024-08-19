package db

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	linkedlist "github.com/PepsiKingIV/KeyValueDB/pkg/db/linked_list"
)

type Store interface {
	Put(string, string) error
	Get(string) (string, error)
	Delete(string) error
	Hash(string, int) (int, error)
	Lock() error
	Unlock() error
	RLock() error
	RUnlock() error
}

type DB struct {
	mu                        sync.RWMutex
	dbSize                    int
	arrayOfPointers           []*linkedlist.Node
	expandTriggerThreshold    float32
	reductionTriggerThreshold float32
	expansionFactors          []float32
	expansionNumber           int
	recordCount               float32
	runtimeOn                 bool
}

func NewDB(ctx context.Context, runtimeOn bool) *DB {
	db := DB{
		mu:                        sync.RWMutex{},
		dbSize:                    100,
		arrayOfPointers:           make([]*linkedlist.Node, 100),
		expandTriggerThreshold:    0.5,
		reductionTriggerThreshold: 0.05,
		expansionFactors:          []float32{10, 5, 2, 1.5},
		expansionNumber:           0,
		recordCount:               0,
		runtimeOn:                 runtimeOn,
	}
	for i := range db.arrayOfPointers {
		db.arrayOfPointers[i] = linkedlist.NewLinkedList()
	}
	if runtimeOn {
		go db.runtime(ctx)
	}
	return &db
}

func (db *DB) expansion() error {
	var expansionNumber int = 3
	if db.expansionNumber < 3 {
		expansionNumber = db.expansionNumber
	}
	newDBSize := int(float32(db.dbSize) * db.expansionFactors[expansionNumber])
	newArrayOfPointers := make([]*linkedlist.Node, newDBSize)
	for i := range newDBSize {
		newArrayOfPointers[i] = linkedlist.NewLinkedList()
	}
	db.RLock()
	for i := range db.arrayOfPointers {
		node := db.arrayOfPointers[i].NextNode
		if node == nil {
			continue
		}
		for node.NextNode != nil {
			ind, err := db.Hash(node.Key, newDBSize)
			if err != nil {
				return errors.Join(errors.New("table expansion error"), err)
			}
			linkedlist.Add(newArrayOfPointers[ind], node.Key, node.Value)
			node = node.NextNode
		}
	}
	db.arrayOfPointers = newArrayOfPointers
	db.dbSize = newDBSize
	db.expansionNumber += 1
	db.RUnlock()
	return nil
}

func (db *DB) reducing() error {
	var expansionNumber int = 3
	if db.expansionNumber < 3 {
		expansionNumber = db.expansionNumber
	}
	newDBSize := db.dbSize / int(db.expansionFactors[expansionNumber])
	newArrayOfPointers := make([]*linkedlist.Node, newDBSize)
	for i := range newDBSize {
		newArrayOfPointers[i] = linkedlist.NewLinkedList()
	}
	db.RLock()
	for i := range db.arrayOfPointers {
		node := db.arrayOfPointers[i].NextNode
		if node == nil {
			continue
		}
		for node.NextNode != nil {
			ind, err := db.Hash(node.Key, newDBSize)
			if err != nil {
				return errors.Join(errors.New("table expansion error"), err)
			}
			linkedlist.Add(newArrayOfPointers[ind], node.Key, node.Value)
			node = node.NextNode
		}
	}
	db.arrayOfPointers = newArrayOfPointers
	db.dbSize = newDBSize
	db.expansionNumber -= 1
	db.RUnlock()
	return nil
}

func (db *DB) runtime(ctx context.Context) {
	for {
		if db.recordCount/float32(db.dbSize) > db.expandTriggerThreshold {
			err := db.expansion()
			fmt.Println("db expantion")
			if err != nil {
				panic(err)
			}
		} else if db.recordCount/float32(db.dbSize) < db.reductionTriggerThreshold && db.expansionNumber > 0 {
			fmt.Println("db reducing")
			err := db.reducing()
			if err != nil {
				panic(err)
			}
		}
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(db.dbSize, "\t", db.recordCount)
			time.Sleep(5 * time.Second)
		}
	}
}

func (db *DB) Hash(key string, dbSize int) (int, error) {
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
	return sum % dbSize, nil
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
	db.Lock()
	defer db.Unlock()
	ind, err := db.Hash(key, db.dbSize)
	if err != nil {
		return ErrInvalidKey
	}
	linkedlist.Add(db.arrayOfPointers[ind], key, value)
	db.recordCount += 1
	return nil
}

// TODO: добавить передачу по ссылке переменной для записи value
func (db *DB) Get(key string) (string, error) {
	switch {
	case (len(key) < 1):
		return "", ErrInvalidKey
	}
	db.RLock()
	defer db.RUnlock()
	ind, err := db.Hash(key, db.dbSize)
	if err != nil {
		return "", ErrInvalidKey
	}
	rec, err := linkedlist.Get(db.arrayOfPointers[ind], key)
	switch {
	case (err == linkedlist.ErrNotFound):
		return "", ErrRecordNotFound
	case (err != nil):
		return "", err
	}
	return rec.Value, err
}

func (db *DB) Delete(key string) error {
	switch {
	case (len(key) < 1):
		return ErrInvalidKey
	}
	db.Lock()
	defer db.Unlock()
	ind, err := db.Hash(key, db.dbSize)
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
	db.recordCount -= 1
	return nil
}
