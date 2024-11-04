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
	softDelete     *Pair
	m              sync.RWMutex
}

func (a *ArrayHashMap) Set(key int, value string) {
	a.m.Lock()
	defer a.m.Unlock()
	index := a.Find(key)
	a.buckets[index] = &Pair{key, value}
}

func (a *ArrayHashMap) Find(key int) int {
	index := a.HasFunc(key)
	findSortDeleteIndex := -1
	for a.buckets[index] != nil {
		if a.buckets[index].key == key {
			if findSortDeleteIndex != -1 {
				a.buckets[findSortDeleteIndex] = a.buckets[index]
				a.buckets[index] = a.softDelete
				return findSortDeleteIndex
			}
			return index
		}
		if findSortDeleteIndex == -1 && a.buckets[index] == a.softDelete {
			findSortDeleteIndex = index
		}
		index = (index + 1) % a.bucketCapacity
	}
	if findSortDeleteIndex != -1 {
		return findSortDeleteIndex
	}
	return index
}

func (a *ArrayHashMap) Get(key int) (v string, ok bool) {
	a.m.RLock()
	defer a.m.RUnlock()
	a.Find(key)
	return "", false
}

func (a *ArrayHashMap) Del(key int) {
	a.m.Lock()
	defer a.m.Unlock()
	index := a.Find(key)
	a.buckets[index] = a.softDelete
}

func (a *ArrayHashMap) HasFunc(key int) int {
	return key % a.bucketCapacity
}

func (a *ArrayHashMap) Print() {
	fmt.Println("------print start--------")
	for _, pair := range a.buckets {
		if pair != nil && pair != a.softDelete {
			fmt.Println(pair.key, pair.value)
		}
	}
	fmt.Println("------print end--------")
}

func (a *ArrayHashMap) Keys() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil && pair != a.softDelete {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

func (a *ArrayHashMap) Values() []string {
	var values []string
	for _, pair := range a.buckets {
		if pair != nil && pair != a.softDelete {
			values = append(values, pair.value)
		}
	}
	return values
}

func NewArrayHashMap(bucketCapacity int) *ArrayHashMap {
	return &ArrayHashMap{
		bucketSize:     0,
		bucketCapacity: bucketCapacity,
		buckets:        make([]*Pair, bucketCapacity),
		softDelete:     &Pair{-1, "-1"},
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
