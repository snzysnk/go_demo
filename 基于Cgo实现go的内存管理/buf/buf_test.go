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
