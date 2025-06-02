package errors

import (
	"fmt"

	"github.com/li-zeyuan/go-leveldb/storage"
)

type ErrCorrupted struct {
	Fd  storage.FileDesc
	Err error
}

func (e *ErrCorrupted) Error() string {
	if !e.Fd.Zero() {
		return fmt.Sprintf("%v [file=%v]", e.Err, e.Fd)
	}

	return e.Err.Error()
}

func NewErrCorrupted(fd storage.FileDesc, err error) error {
	return &ErrCorrupted{
		Fd:  fd,
		Err: err,
	}
}
