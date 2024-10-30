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
	buckets        []*Pair
	m              sync.RWMutex
}

func (a ArrayHashMap) Set(key int, value string) {
	a.m.Lock()
	defer a.m.Unlock()
	index := a.HasFunc(key)
	a.buckets[index] = &Pair{key, value}
}

func (a ArrayHashMap) Get(key int) (v string, ok bool) {
	a.m.RLock()
	defer a.m.RUnlock()
	index := a.HasFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return "", false
	}
	return pair.value, true
}

func (a ArrayHashMap) Del(key int) {
	a.m.Lock()
	defer a.m.Unlock()
	a.buckets[a.HasFunc(key)] = nil
}

func (a ArrayHashMap) HasFunc(key int) int {
	return key % a.bucketCapacity
}

func (a ArrayHashMap) Print() {
	fmt.Println("------print start--------")
	for _, pair := range a.buckets {
		if pair == nil {
			continue
		}
		fmt.Println(pair.key, "->", pair.value)
	}
	fmt.Println("------print end--------")
}

func (a ArrayHashMap) Keys() []int {
	keys := make([]int, 0)
	for _, pair := range a.buckets {
		if pair == nil {
			continue
		}
		keys = append(keys, pair.key)
	}
	return keys
}

func (a ArrayHashMap) Values() []string {
	values := make([]string, 0)
	for _, pair := range a.buckets {
		if pair == nil {
			continue
		}
		values = append(values, pair.value)
	}
	return values
}

func NewArrayHashMap(bucketCapacity int) *ArrayHashMap {
	return &ArrayHashMap{
		bucketSize:     0,
		bucketCapacity: bucketCapacity,
		buckets:        make([]*Pair, bucketCapacity),
	}
}

//go:generate go run main.go
func main() {
	hashMap := NewArrayHashMap(100)
	hashMap.Set(0, "00")
	hashMap.Set(1, "01")
	hashMap.Set(2, "02")
	hashMap.Print()
	hashMap.Del(1)
	hashMap.Print()
	fmt.Println(hashMap.Keys())
	fmt.Println(hashMap.Values())
}
