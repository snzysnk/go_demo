package validate

import "fmt"

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
