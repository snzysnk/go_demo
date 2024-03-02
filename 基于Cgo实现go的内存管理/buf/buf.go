package buf

import (
	"go_demo/基于Cgo实现go的内存管理/c"
	"unsafe"
)

type IBuf interface {
	SetBytes([]byte)
	GetBytes() []byte
	Free()
	Copy(buf *IBuf)
}

type Buf struct {
	Next     *Buf
	Capacity int
	length   int
	head     int
	data     unsafe.Pointer
}

func (b *Buf) Copy(buf *IBuf) {
	//TODO implement me
	panic("implement me")
}

func (b *Buf) Free() {
	c.Free(b.data)
}

func (b *Buf) SetBytes(bytes []byte) {
	c.MemCopy(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), bytes, len(bytes))
	b.length += len(bytes)
}

func (b *Buf) GetBytes() []byte {
	return c.GetBytes(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), b.length)
}

func NewBuf(size int) *Buf {
	return &Buf{
		Next:     nil,
		Capacity: size,
		length:   0,
		head:     0,
		data:     c.Malloc(size),
	}
}
