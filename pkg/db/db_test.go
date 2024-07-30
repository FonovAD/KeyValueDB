package db_test

import (
	"os"
	"testing"

	"github.com/PepsiKingIV/KeyValueDB/pkg/db"
	DB "github.com/PepsiKingIV/KeyValueDB/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
func TestHash(t *testing.T) {
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
			Expected: DB.ErrInvalidKey,
		},
		TestCase{
			ID:       4,
			Name:     "Too long key",
			Key:      "qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwe",
			dbSize:   100,
			Expected: DB.ErrKeyTooLong,
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
		db := DB.NewDB()

		a, err := db.Hash(tc.Key)
		assert.Equal(t, tc.Expected, err)
		b, _ := db.Hash(tc.Key)
		assert.Equal(t, a, b)
	}
}

func TestLock(t *testing.T) {
	db := DB.NewDB()

	err := db.Lock()
	db.Unlock()
	assert.NoError(t, err)
}

func TestUnlock(t *testing.T) {
	db := DB.NewDB()

	db.Lock()
	err := db.Unlock()
	assert.NoError(t, err)
}

func TestRLock(t *testing.T) {
	db := DB.NewDB()

	err := db.RLock()
	db.RUnlock()
	assert.NoError(t, err)
}

func TestRUnlock(t *testing.T) {
	db := DB.NewDB()

	db.RLock()
	err := db.RUnlock()
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
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
				db := DB.NewDB()

				err := db.Put("Key", "Value")
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
			Name:     "Empty Key",
			Prepare:  func() {},
			Key:      "",
			Value:    "Value",
			Expected: db.ErrInvalidKey,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			db := DB.NewDB()

			if len(tc.Key) != 0 {
				//плохо, что результаты теста зависят от db.Put
				err := db.Put(tc.Key, tc.Value)
				assert.NoError(t, err)
			}

			tc.Prepare()
			value, err := db.Get(tc.Key)
			assert.Equal(t, tc.Expected, err)
			if err == nil {
				assert.Equal(t, tc.Value, value)
			}
		})
	}
}

func TestPut(t *testing.T) {
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
		db := DB.NewDB()

		err := db.Put(tc.Key, tc.Value)
		assert.Equal(t, tc.Expected, err)
	}
}

func TestDelete(t *testing.T) {
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
			Name:     "record not found",
			Key:      "Key",
			Value:    "Value",
			Expected: db.ErrRecordNotFound,
		},
	}
	for _, tc := range tcs {
		db := DB.NewDB()

		err := db.Put(tc.Key, tc.Value)
		assert.NoError(t, err)

		err = db.Delete(tc.Key)
		assert.Equal(t, tc.Expected, err)
	}
}
