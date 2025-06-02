package storage

import "io"

type FileType int

const (
	TypeManifest FileType = 1 << iota
	TypeJournal
	TypeTable
	TypeTemp

	TypeAll = TypeManifest | TypeJournal | TypeTable | TypeTemp
)

type Storage interface {
	Lock() (Lock, error)
	Log(str string)
	SetMeta(fd FileDesc) error
	GetMeta() (FileDesc, error)
	List(ft FileType) ([]FileDesc, error)
	Open(fd FileDesc) (Reader, error)
	Create(fd FileDesc) (Writer, error)
	Remove(fd FileDesc) error
	Rename(olddf, newfd FileDesc) error
	Close() error
}

type FileDesc struct {
	Type FileType
	Num  int64
}

func (fd FileDesc) Zero()bool {
	return fd == FileDesc{}
}

type Lock interface {
	Unlock()
}

type Reader interface {
	io.ReadSeeker
	io.ReaderAt
	io.Closer
}

type Writer interface {
	io.WriteCloser
	Syncer
}

type Syncer interface {
	Sync() error
}


