package sl

import (
	"errors"
	"read_file_config/lk"
	"sync"
)

type IXStrSlice interface {
	lk.IXLock
	Set(index int, value string) error
	Get(index int) (value string, found bool)
	Add(item string)
}

var _ IXStrSlice = (*XStrSlice)(nil)

type XStrSlice struct {
	lk.IXLock
	data []string
}

func (x XStrSlice) Add(item string) {
	x.Lock()
	defer x.Unlock()
	x.data = append(x.data, item)
}

func NewSafeXStrSlice(data []string) IXStrSlice {
	return XStrSlice{
		IXLock: new(sync.RWMutex),
		data:   data,
	}
}

func NewUnSafeXStrSlice(data []string) IXStrSlice {
	return XStrSlice{
		IXLock: new(lk.XLock),
		data:   data,
	}
}

func (x XStrSlice) Set(index int, value string) error {
	x.Lock()
	defer x.Unlock()
	if index < 0 || index >= len(x.data) {
		return errors.New("越界了")
	}
	x.data[index] = value
	return nil
}

func (x XStrSlice) Get(index int) (value string, found bool) {
	x.Lock()
	defer x.Unlock()
	if index < 0 || index >= len(x.data) {
		return "", false
	}
	return x.data[index], true
}
