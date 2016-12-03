package pool

import (
	"errors"
)

var (
	ErrClosed = errors.New("对象池已关闭")
)

type Pool interface {
	Get() ([]byte, error)
	Put([]byte) error
	Close([]byte) error
	Release()
	Len() int
}
