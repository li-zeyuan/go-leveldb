package storage

import (
	"log"
	"os"
	"syscall"
)

type unixFileLock struct {
	f *os.File
}

func (fl *unixFileLock) release() error {
	if err := setFileLock(fl.f, false, false); err != nil {
		return err
	}
	
	return fl.f.Close()
}

func newFileLock(path string, readOnly bool) (fileLock, error) {
	var flag int
	if readOnly {
		flag = os.O_RDONLY
	} else {
		flag = os.O_CREATE | os.O_RDWR
	}
	f, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		log.Printf("leveldb/storage: open %s: %v", path, err)
		return nil, err
	}

	if err = setFileLock(f, readOnly, true); err != nil {
		f.Close()
		return nil, err
	}
	return &unixFileLock{f: f}, nil
}

func setFileLock(f *os.File, readOnly bool, lock bool) error {
	how := syscall.LOCK_UN
	if lock {
		if readOnly {
			// shared lock
			how = syscall.LOCK_SH
		} else {
			// exclusive lock
			how = syscall.LOCK_EX
		}
	}

	// non-blocking
	return syscall.Flock(int(f.Fd()), how|syscall.LOCK_NB)
}
