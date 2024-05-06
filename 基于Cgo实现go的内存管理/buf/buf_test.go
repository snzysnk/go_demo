package buf

import (
	"fmt"
	"testing"
)

func TestBuf(t *testing.T) {
	buf := NewBuf(4)
	buf.SetBytes([]byte("123"))
	fmt.Println(string(buf.GetBytes()))
	buf.Free()
}

func TestBufCopy(t *testing.T) {
	var (
		buf01 = NewBuf(4)
		buf02 = NewBuf(4)
	)
	defer buf01.Free()
	defer buf02.Free()
	buf01.SetBytes([]byte("123"))
	buf02.Copy(buf01)
	fmt.Println(string(buf01.GetBytes()))
	fmt.Println(string(buf02.GetBytes()))
}
