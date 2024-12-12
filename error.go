package goparse

import (
	"errors"
	"fmt"
)

var (
	ERR_CREATE         = errors.New("goparse: error create template")
	ERR_CREATE_TMP_DIR = errors.New("failed created temporary directory")
)

func ReturnPanic(err error) {
	panic(errors.New(fmt.Sprintf("%s\n%s", ERR_CREATE.Error(), err.Error())))
}

func ErrDirDoesntExist(dir string) error {
	return errors.New(fmt.Sprintf("%s\ndirectory %s does not exists", ERR_CREATE.Error(), dir))
}
