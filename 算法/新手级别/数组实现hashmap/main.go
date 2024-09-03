package main

import (
	"fmt"
	"sync"
)

type Pair struct {
	key   int
	value interface{}
}

type HashMap interface {
	Set(key int, value interface{})
	Get(key int) (v interface{}, ok bool)
	Del(key int)
	loadFactor() float64
	extend()
}

type ArrayHashMap struct {
	bucketSize     int
	bucketCapacity int
	buckets        [][]Pair
	loadThres      float64
	extendRation   int
	m              sync.RWMutex
}

func NewArrayHashMap(bucketCapacity int) *ArrayHashMap {
	return &ArrayHashMap{
		bucketSize:     0,
		bucketCapacity: bucketCapacity,
		buckets:        make([][]Pair, bucketCapacity),
		loadThres:      2.0 / 3.0,
		extendRation:   2,
	}
}

func (a *ArrayHashMap) Set(key int, value interface{}) {
	a.m.Lock()
	defer a.m.Unlock()
	index := a.HashFunc(key)
	for i, p := range a.buckets[index] {
		if p.key == key {
			a.buckets[index][i].value = value
			return
		}
	}
	a.buckets[index] = append(a.buckets[index], Pair{key, value})
	a.bucketSize++
}

func (a *ArrayHashMap) Get(key int) (v interface{}, ok bool) {
	a.m.RLock()
	defer a.m.RUnlock()
	index := a.HashFunc(key)
	for _, p := range a.buckets[index] {
		if p.key == key {
			return p.value, true
		}
	}
	return nil, false
}

func (a *ArrayHashMap) Del(key int) {
	a.m.Lock()
	defer a.m.Unlock()
	index := a.HashFunc(key)
	for i, p := range a.buckets[index] {
		if p.key == key {
			a.bucketSize--
			a.buckets[index] = append(a.buckets[index][:i], a.buckets[index][i+1:]...)
			return
		}
	}
}

func (a *ArrayHashMap) dump() {
	var sl []interface{}
	for _, p := range a.buckets {
		for _, pair := range p {
			sl = append(sl, pair.value)
		}
	}
	fmt.Print(sl)
}

func (a *ArrayHashMap) loadFactor() float64 {
	return float64(a.bucketSize) / float64(a.bucketCapacity)
}

func (a *ArrayHashMap) extend() {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayHashMap) HashFunc(key int) int {
	return key % a.bucketCapacity
}

var _ HashMap = (*ArrayHashMap)(nil)

//go:generate go run main.go
func main() {
	hashMap := NewArrayHashMap(10)
	hashMap.Set(1, 101)
	hashMap.dump()
	if v, ok := hashMap.Get(1); ok {
		fmt.Println(v)
	}
	hashMap.Del(1)
	hashMap.dump()
}
