package db

type Store interface {
	Put(string, string) error
	Get(string) (string, error)
	Delete(string) error
	Lock() error
	Unlock() error
	RLock() error
	RUnlock() error
}

const (
	Unlock = 0
	Lock   = 1
	RLock  = 2
)

type DB struct {
	lock            int
	dbSize          int
	arrayOfPointers [][]string
}

func NewDB() DB {
	return DB{
		lock:            Unlock,
		dbSize:          100,
		arrayOfPointers: [][]string{},
	}
}
