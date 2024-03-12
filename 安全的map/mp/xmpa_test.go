package mp

import (
	"strconv"
	"sync"
	"testing"
)

func TestSafeMap(t *testing.T) {
	testXMap(CreateSafeXMap())
}

func TestUnSafeMap(t *testing.T) {
	//竞态
	testXMap(CreateUnSafeXMap())
}

func testXMap(m IXMap) {
	var (
		group sync.WaitGroup
	)
	group.Add(100)
	for i := 0; i < 100; i++ {
		value := strconv.Itoa(i)
		key := strconv.Itoa(i % 2)
		go func() {
			m.Set(key, value)
			group.Done()
		}()
	}
	group.Wait()
}
