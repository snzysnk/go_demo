package main

import (
	"fmt"
	"strconv"
	"strings"
)

// PrintBinary 打印二进制
func PrintBinary(num int) string {
	var result string
	for i := 31; i >= 0; i-- {
		if num&(1<<i) == 0 {
			result = result + "0"
		} else {
			result = result + "1"
		}
	}
	return result
}

// RestoreBinary 还原二进制字符串
func RestoreBinary(str string) int {
	var (
		result           int
		strSlice         = strings.Split(str, "")
		isPositiveNumber = string(strSlice[0]) == "0"
	)

	for i := 1; i <= 31; i++ {
		n, _ := strconv.Atoi(strSlice[i])
		if !isPositiveNumber {
			if n == 0 {
				n = 1
			} else {
				n = 0
			}
		}
		result += (1 << (31 - i)) * n
	}

	if isPositiveNumber {
		return result
	}

	return -(result + 1)
}

func main() {
	fmt.Println(PrintBinary(2))  //00000000000000000000000000000010
	fmt.Println(PrintBinary(4))  //00000000000000000000000000000100
	fmt.Println(PrintBinary(7))  //00000000000000000000000000000111
	fmt.Println(PrintBinary(10)) //00000000000000000000000000001010
	fmt.Println(PrintBinary(-1)) //11111111111111111111111111111111
	fmt.Println(PrintBinary(-2)) //11111111111111111111111111111110

	fmt.Println(RestoreBinary(PrintBinary(2)))    //2
	fmt.Println(RestoreBinary(PrintBinary(10)))   //10
	fmt.Println(RestoreBinary(PrintBinary(-101))) //101
}
