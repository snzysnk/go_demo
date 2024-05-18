package main

import "fmt"

var (
	example = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	}
	insertPosition = 4
	insertValue    = "x"
)

func First() {
	result := make([]string, 0)
	result = append(result, example[:insertPosition]...)
	result = append(result, insertValue)
	result = append(result, example[insertPosition:]...)
	fmt.Println(result)
}

func Next() {
	example = append(example, "")

	//底层array a b c d e  可见部分 f g h i j k l m  ""
	// a b c d e e f g h i j k l m
	copy(example[insertPosition+1:], example[insertPosition:])
	example[insertPosition] = insertValue
	fmt.Println(example)
}

func Three() {
	insertPosition = 0
	example[insertPosition] = "55555"
	example = append(example, "")
	copy(example[insertPosition+2:], example[insertPosition+1:])
	example[insertPosition+1] = "6666"
	fmt.Println(example)
}

func main() {
	First() //[a b c d x e f g h i j k l m]
	Next()  //[a b c d x e f g h i j k l m]
	Three()
}
