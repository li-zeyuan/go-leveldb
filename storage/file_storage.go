package storage

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type fileLock interface {
	release() error
}

type fileStorage struct {
	path     string
	readOnly bool
	flock    fileLock
	logw     *os.File
	logSize  int64
}

func OpenFile(path string, readOnly bool) (Storage, error) {
	// path is a directory
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) && !readOnly {
			if err := os.MkdirAll(path, 0755); err != nil {
				log.Printf("leveldb/storage: open %s: %v", path, err)
				return nil, err
			}
		} else {
			log.Printf("leveldb/storage: open %s: %v", path, err)
			return nil, err
		}
	}
	if !fi.IsDir() {
		log.Printf("leveldb/storage: open %s: not a directory", path)
		return nil, fmt.Errorf("leveldb/storage: open %s: not a directory", path)
	}

	flock, err := newFileLock(filepath.Join(path, "LOCK"), readOnly)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			flock.release()
		}
	}()

	var logw *os.File
	var logSize int64

	if !readOnly {
		logw, err = os.OpenFile(filepath.Join(path, "LOG"), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}

		logSize, err = logw.Seek(0, io.SeekEnd)
		if err != nil {
			logw.Close()
			return nil, err
		}
	}

	fs := &fileStorage{
		path:     path,
		readOnly: readOnly,
		flock:    flock,
		logw:     logw,
		logSize:  logSize,
	}

	// gc will call this method
	runtime.SetFinalizer(fs, (*fileStorage).Close)
	return fs, nil
}

// todo 实现Storage接口

// todo
func (fs *fileStorage) Close() error {
	return nil
}
