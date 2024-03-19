package sl

import (
	"strconv"
	"sync"
	"testing"
)

func TestSafeStrSlice(t *testing.T) {
	testXStrSlice(NewSafeXStrSlice([]string{"a", "b"}))
}

func TestUnSafeStrSlice(t *testing.T) {
	testXStrSlice(NewUnSafeXStrSlice([]string{"a", "b"}))
}

func testXStrSlice(m IXStrSlice) {
	var (
		group sync.WaitGroup
	)
	group.Add(100)
	for i := 0; i < 100; i++ {
		key := i % 2
		value := i
		go func() {
			_ = m.Set(key, strconv.Itoa(value))
			group.Done()
		}()
	}
	group.Wait()
}
