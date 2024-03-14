package mp

import (
	"go_demo/安全的map或切片/lk"
	"sync"
)

type IXMap interface {
	lk.IXLock
	Set(key, value string)
	Get(key string) string
}

var _ IXMap = (*XMap)(nil)

type XMap struct {
	lk.IXLock
	data map[string]string
}

func (x XMap) Set(key, value string) {
	x.Lock()
	defer x.Unlock()
	x.data[key] = value
}

func (x XMap) Get(key string) string {
	x.Lock()
	defer x.Unlock()
	return x.data[key]
}

func CreateSafeXMap() IXMap {
	return XMap{
		IXLock: &sync.RWMutex{},
		data:   make(map[string]string),
	}
}

func CreateUnSafeXMap() IXMap {
	return XMap{
		IXLock: &lk.XLock{},
		data:   make(map[string]string),
	}
}
