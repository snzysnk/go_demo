package main

import "fmt"

//go:generate go run main.go
func main() {
	fmt.Println(itoa(0, 3))    //0000
	fmt.Println(itoa(2024, 3)) //2024
	fmt.Println(itoa(202, 5))  //00202
}

// 该方法改编自log标准包
func itoa(i int, wid int) string {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	return string(b[bp:])
}
