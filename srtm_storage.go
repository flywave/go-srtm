package srtm

import (
	"io/ioutil"
	"os"
	"path"
)

type SrtmStorage interface {
	LoadFile(fn string) ([]byte, error)
	IsNotExists(err error) bool
	SaveFile(fn string, bytes []byte) error
}

type LocalFileSrtmStorage struct {
	cacheDirectory string
}

func NewLocalFileSrtmStorage(cacheDirectory string) (*LocalFileSrtmStorage, error) {
	if len(cacheDirectory) == 0 {
		cacheDirectory = path.Join(os.Getenv("HOME"), ".srtm")
	}

	if _, err := os.Stat(cacheDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(cacheDirectory, os.ModeDir|0700); err != nil {
			return nil, err
		}
	}

	return &LocalFileSrtmStorage{cacheDirectory: cacheDirectory}, nil
}

func (ds LocalFileSrtmStorage) LoadFile(fn string) ([]byte, error) {
	f, err := os.Open(path.Join(ds.cacheDirectory, fn))
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (ds LocalFileSrtmStorage) IsNotExists(err error) bool {
	return os.IsNotExist(err)
}

func (ds LocalFileSrtmStorage) SaveFile(fn string, bytes []byte) error {
	f, err := os.Create(path.Join(ds.cacheDirectory, fn))
	if err != nil {
		return err
	}
	_, err = f.Write(bytes)
	return err
}

var _ SrtmStorage = new(LocalFileSrtmStorage)
