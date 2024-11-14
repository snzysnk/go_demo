package main

import (
	"fmt"
	"sync"
)

type Pair struct {
	key   int
	value string
}

var _ HashMap = (*ArrayHashMap)(nil)

type HashMap interface {
	Set(key int, value string)
	Get(key int) (v string, ok bool)
	Del(key int)
	HasFunc(key int) int
	Print()
	Keys() []int
	Values() []string
}

type ArrayHashMap struct {
	bucketSize     int
	bucketCapacity int
	buckets        [][]*Pair
	m              sync.RWMutex
}

func (a *ArrayHashMap) Set(key int, value string) {
	a.m.Lock()
	defer a.m.Unlock()
	index := a.HasFunc(key)
	origIndex := a.Find(index, key)
	//存在就更新
	if origIndex >= 0 {
		a.buckets[index][origIndex] = &Pair{key, value}
		return
	}
	a.buckets[index] = append(a.buckets[index], &Pair{key, value})
}

func (a *ArrayHashMap) Find(hashIndex, key int) int {
	pairs := a.buckets[hashIndex]
	for index, pair := range pairs {
		if pair.key == key {
			return index
		}
	}
	return -1
}

func (a *ArrayHashMap) Get(key int) (v string, ok bool) {
	a.m.RLock()
	defer a.m.RUnlock()
	index := a.HasFunc(key)
	pairs := a.buckets[index]
	for _, pair := range pairs {
		if pair.key == key {
			return pair.value, true
		}
	}
	return "", false
}

func (a *ArrayHashMap) Del(key int) {
	a.m.Lock()
	defer a.m.Unlock()
	index := a.HasFunc(key)
	pairs := a.buckets[index]
	for i, pair := range pairs {
		if pair.key == key {
			a.buckets[index] = append(pairs[:i], pairs[i+1:]...)
		}
	}
}

func (a *ArrayHashMap) HasFunc(key int) int {
	return key % a.bucketCapacity
}

func (a *ArrayHashMap) Print() {
	fmt.Println("------print start--------")
	for _, pairs := range a.buckets {
		for _, pair := range pairs {
			if pair == nil {
				continue
			}
			fmt.Println(pair.key, "->", pair.value)
		}
	}
	fmt.Println("------print end--------")
}

func (a *ArrayHashMap) Keys() []int {
	keys := make([]int, 0)
	for _, pairs := range a.buckets {
		for _, pair := range pairs {
			if pair == nil {
				continue
			}
			keys = append(keys, pair.key)
		}

	}
	return keys
}

func (a *ArrayHashMap) Values() []string {
	values := make([]string, 0)
	for _, pairs := range a.buckets {
		for _, pair := range pairs {
			if pair == nil {
				continue
			}
			values = append(values, pair.value)
		}
	}
	return values
}

func NewArrayHashMap(bucketCapacity int) *ArrayHashMap {
	return &ArrayHashMap{
		bucketSize:     0,
		bucketCapacity: bucketCapacity,
		buckets:        make([][]*Pair, bucketCapacity),
	}
}

//go:generate go run main.go
func main() {
	hashMap := NewArrayHashMap(100)
	hashMap.Set(0, "00")
	hashMap.Set(1, "01")
	hashMap.Set(101, "101")
	hashMap.Set(2, "02")
	hashMap.Print()
	hashMap.Del(1)
	hashMap.Set(101, "101_UPDATE")
	hashMap.Print()
	fmt.Println(hashMap.Keys())
	fmt.Println(hashMap.Values())
}
