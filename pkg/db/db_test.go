package db_test

import (
	"testing"

	"github.com/PepsiKingIV/KeyValueDB/pkg/db"
	"github.com/stretchr/testify/assert"
)

func HashTest(t *testing.T) {
	type TestCase struct {
		ID       int
		Name     string
		Key      string
		dbSize   int
		Expected error
	}
	tcs := []TestCase{
		TestCase{
			ID:       1,
			Name:     "basic case, db_size = 100",
			Key:      "FirstTest",
			dbSize:   100,
			Expected: nil,
		},
		TestCase{
			ID:       2,
			Name:     "basic case, db_size = 1000",
			Key:      "SecondTest",
			dbSize:   1000,
			Expected: nil,
		},
		TestCase{
			ID:       3,
			Name:     "Empty key",
			Key:      "",
			dbSize:   100,
			Expected: db.ErrInvalidKey,
		},
		TestCase{
			ID:       4,
			Name:     "Too long key",
			Key:      "qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwe",
			dbSize:   100,
			Expected: db.ErrKeyTooLong,
		},
		TestCase{
			ID:       5,
			Name:     "Only numbers",
			Key:      "124189120",
			dbSize:   100,
			Expected: nil,
		},
	}
	for _, tc := range tcs {
		a, err := Hash(tc.Key, tc.dbSize)
		b := Hash(tc.Key, tc.dbSize)
		assert.Equal(t, tc.Expected, err)
		assert.Equal(t, a, b)
	}
}

func LockTest(t *testing.T) {
	err := Lock()
	assert.NoError(t, err)
}

func UnlockTest(t *testing.T) {
	err := Unlock()
	assert.NoError(t, err)
}

func RLockTest(t *testing.T) {
	err := RLock()
	assert.NoError(t, err)
}

func RUnlockTest(t *testing.T) {
	err := RUnlock()
	assert.NoError(t, err)
}

func GetTest(t *testing.T) {
	type TestCase struct {
		ID       int
		Name     string
		Prepare  func()
		Key      string
		Value    string
		Expected error
	}
	tcs := []TestCase{
		TestCase{
			ID:   1,
			Name: "Basic case",
			Prepare: func() {
				db, err := NewDB()
				assert.NoError(t, err)

				err = db.Put("Key", "Value")
				assert.NoError(t, err)
			},
			Key:   "Key",
			Value: "Value",
		},
		TestCase{
			ID:       1,
			Name:     "Basic case",
			Prepare:  func() {},
			Key:      "Key",
			Value:    "Value",
			Expected: nil,
		},
		TestCase{
			ID:       2,
			Name:     "Empty Value",
			Prepare:  func() {},
			Key:      "Key",
			Value:    "",
			Expected: db.ErrInvalidValue,
		},
		TestCase{
			ID:       2,
			Name:     "Empty Key",
			Prepare:  func() {},
			Key:      "",
			Value:    "Value",
			Expected: db.ErrInvalidKey,
		},
	}
	for _, tc := range tcs {
		db, err := NewDB()
		assert.NoError(t, err)

		//плохо, что результаты теста зависят от db.Put
		err = db.Put(tc.Key, tc.Value)
		assert.NoError(t, err)

		tc.Prepare()
		value, err := db.Get(tc.Key)
		assert.Equal(t, tc.Expected, err)
		assert.Equal(t, tc.Value, value)
	}
}

func PutTest(t *testing.T) {
	type TestCase struct {
		ID       int
		Name     string
		Key      string
		Value    string
		Expected error
	}
	tcs := []TestCase{
		TestCase{
			ID:       1,
			Name:     "Basic case",
			Key:      "Key",
			Value:    "Value",
			Expected: nil,
		},
		TestCase{
			ID:       2,
			Name:     "Empty key",
			Key:      "",
			Value:    "Value",
			Expected: db.ErrInvalidKey,
		},
		TestCase{
			ID:       3,
			Name:     "Empty value",
			Key:      "Key",
			Value:    "",
			Expected: db.ErrInvalidValue,
		},
	}
	for _, tc := range tcs {
		db, err := NewDB()
		assert.NoError(t, err)

		err = db.Put(tc.Key, tc.Value)
		assert.Equal(t, tc.Expected, err)
	}
}
