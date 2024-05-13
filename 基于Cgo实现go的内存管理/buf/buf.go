package buf

import (
	"go_demo/c"
	"unsafe"
)

type IBuf interface {
	SetBytes([]byte)  //设置数据
	GetBytes() []byte //获取数据
	Free()            //释放数据
	Copy(buf IBuf)    //拷贝数据
	GetLength() int
}

var _ IBuf = (*Buf)(nil)

type Buf struct {
	Next     *Buf
	Capacity int
	length   int
	head     int
	data     unsafe.Pointer
}

func NewBuf(size int) IBuf {
	return &Buf{
		length: 0,
		head:   0,
		data:   c.Malloc(size),
	}
}

func (b *Buf) Copy(other IBuf) {
	c.MemCopy(b.data, other.GetBytes(), other.GetLength())
	b.head = 0
	b.length = other.GetLength()
}

func (b *Buf) Free() {
	c.Free(b.data)
}

func (b *Buf) GetLength() int {
	return b.length
}

func (b *Buf) SetBytes(bytes []byte) {
	c.MemCopy(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), bytes, len(bytes))
	b.length += len(bytes)
}

func (b *Buf) GetBytes() []byte {
	return c.GetBytes(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), b.length)
}
