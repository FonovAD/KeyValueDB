package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func HashTest(t *testing.T) {
	type TestCase struct {
		ID      int
		Name    string
		Key     string
		db_size int
	}
	tcs := []TestCase{
		TestCase{
			ID:      1,
			Name:    "basic case, db_size = 100",
			Key:     "FirstTest",
			db_size: 100,
		},
		TestCase{
			ID:      2,
			Name:    "basic case, db_size = 1000",
			Key:     "SecondTest",
			db_size: 1000,
		},
		TestCase{
			ID:      3,
			Name:    "Empty key",
			Key:     "",
			db_size: 100,
		},
		TestCase{
			ID:      4,
			Name:    "Too long key",
			Key:     "qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./qwertyuiop[]asdfghjkl;'zxcvbnm,./",
			db_size: 100,
		},
		TestCase{
			ID:      5,
			Name:    "Only numbers",
			Key:     "124189120",
			db_size: 100,
		},
	}
	for _, tc := range tcs {
		a := Hash(tc.Key, tc.db_size)
		b := Hash(tc.Key, tc.db_size)
		assert.Equal(t, a, b)
	}
}
