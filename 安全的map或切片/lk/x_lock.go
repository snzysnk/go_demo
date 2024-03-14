package lk

type IXLock interface {
	Lock()
	Unlock()
}

type XLock struct {
}

func (X *XLock) Lock() {
	return
}

func (X *XLock) Unlock() {
	return
}
