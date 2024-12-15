package goparse

import (
	"errors"
	"fmt"
)

var (
	ErrCreate         = errors.New("goparse: error create template")
	ErrCreateTmpDir   = errors.New("failed created temporary directory")
	ErrInvalidPattern = errors.New("invalid pattern")
)

func returnPanic(err error) {
	panic(fmt.Errorf("%s\n%s", ErrCreate.Error(), err.Error()))
}

func errDirDoesntExist(dir string) error {
	return fmt.Errorf("%s\ndirectory %s does not exists", ErrCreate.Error(), dir)
}
