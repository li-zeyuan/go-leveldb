package goleveldb

import (
	"os"
	
	"github.com/li-zeyuan/go-leveldb/storage"
	"github.com/li-zeyuan/go-leveldb/errors"
)

type session struct {
	stor storage.Storage
}

func newSession(storage storage.Storage) (*session, error) {
	return nil, nil
}

func (s *session)recover()(err error) {
	defer func () {
		if os.IsNotExist(err) {
			// rewrite os.ErrNotExist error
			if fds, _ := s.stor.List(storage.TypeAll); len(fds) > 0 {
				err = errors.NewErrCorrupted(storage.FileDesc{}, errors.New("database entry point either missing or corrupted"))
			}
		}
	}()

	fd, err := s.stor.GetMeta()
	if err != nil {
		return
	}
}


func (s *session) close() {
}

func (s *session) release() {
}


