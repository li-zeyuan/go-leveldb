package goleveldb

import "github.com/li-zeyuan/go-leveldb/storage"

type DB struct {
	storage storage.Storage
}

func Open(storage storage.Storage) (db *DB, err error) {
	s, err := newSession(storage)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			s.close()
			s.release()
		}
	}()

	return &DB{
		storage: s,
	}, nil
}

