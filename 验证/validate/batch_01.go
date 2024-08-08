package validate

import (
	"fmt"
	"math"
	"reflect"
)

func F1() {
	var ex []interface{}
	if ex == nil {
		fmt.Println("[]interface{} is nil")
	} else {
		fmt.Println("[]interface{} is not nil")
	}
}

func F2() {
	var ex []string
	if ex == nil {
		fmt.Println("[]string is nil")
	} else {
		fmt.Println("[]string not nil")
	}
}

func F3() {
	var a interface{}
	a = []byte{}
	switch a.(type) {
	case []byte:
		fmt.Println("match []byte")
	}
}

func F4() {
	var a interface{}
	a = []byte{}
	switch a.(type) {
	case []uint8:
		fmt.Println("match []uint8")
	}
}

func F5() {
	ch := make(chan struct{}, 100)
	ch <- struct{}{}
	fmt.Println("len获取的是实际长度:", len(ch))
	fmt.Println("cap才是获取的容量:", cap(ch))
	defer close(ch)
}

func F6() {
	fmt.Println(math.Float32bits(1.00))
}

type mF7 struct {
	reflect.Value
}

func F7() {
	var a interface{}
	var b interface{}
	a = reflect.ValueOf("2")
	b = mF7{}
	if value, ok := a.(reflect.Value); ok {
		fmt.Println("a:", value)
	}

	if value, ok := b.(reflect.Value); ok {
		fmt.Println("b:", value)
	}
}
