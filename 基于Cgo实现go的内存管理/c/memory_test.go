package c

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"
)

func IsLittleEndian() bool {
	n := 0x01020304
	u := unsafe.Pointer(&n)
	pb := (*byte)(u)
	b := *pb //8位
	return b == 0x04
}

func intToByte(n uint32) []byte {
	var (
		bytesBuffer = bytes.NewBuffer([]byte{})
	)
	var order binary.ByteOrder
	if IsLittleEndian() {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}
	_ = binary.Write(bytesBuffer, order, int32(n))
	return bytesBuffer.Bytes()
}

func TestMemory(t *testing.T) {
	data := Malloc(4)
	fmt.Printf("data %+v,%T\n", data, data)
	myData := (*uint32)(data)
	*myData = 6
	fmt.Printf("data %+v,%T\n", *myData, *myData)
	var a uint32 = 101
	MemCopy(data, intToByte(a), 4)
	fmt.Printf("data %+v,%T\n", *myData, *myData)
	Free(data)
}
