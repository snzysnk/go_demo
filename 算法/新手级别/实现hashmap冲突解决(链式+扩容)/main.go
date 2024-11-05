package main

import (
	"fmt"
	"log"
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
	Cap() int
	Del(key int)
	Print()
	Keys() []int
	Values() []string
	hasFunc(key int) int
	loadFactor() float64
	extend()
	setWithoutLock(key int, value string)
}

type ArrayHashMap struct {
	bucketSize     int
	bucketCapacity int
	buckets        [][]*Pair
	m              sync.RWMutex
	loadThres      float64
	extendRatio    int
	size           int
}

func (a *ArrayHashMap) Cap() int {
	return a.bucketCapacity
}

func (a *ArrayHashMap) setWithoutLock(key int, value string) {
	a.coreSet(key, value)
}

func (a *ArrayHashMap) loadFactor() float64 {
	return float64(a.size) / float64(a.bucketCapacity)
}

func (a *ArrayHashMap) extend() {
	log.Println("Triggered expansion start")
	// temporary storage old buckets and extend
	oldBuckets := a.buckets
	a.bucketCapacity *= a.extendRatio
	//there need set size equal zero because setWithoutLock can size++
	a.size = 0
	a.buckets = make([][]*Pair, a.bucketCapacity)
	// after extend reset data
	for _, pairs := range oldBuckets {
		for _, pair := range pairs {
			if pair == nil {
				continue
			}
			a.setWithoutLock(pair.key, pair.value)
		}
	}
	log.Println("Triggered expansion end")
}

func (a *ArrayHashMap) Set(key int, value string) {
	a.m.Lock()
	defer a.m.Unlock()
	a.coreSet(key, value)
}

func (a *ArrayHashMap) coreSet(key int, value string) {
	if a.loadFactor() > a.loadThres {
		fmt.Println(a.size, a.bucketCapacity)
		a.extend()
	}
	index := a.hasFunc(key)
	origIndex := a.Find(index, key)
	// if exist update
	if origIndex >= 0 {
		a.buckets[index][origIndex] = &Pair{key, value}
		return
	}
	a.size++
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
	index := a.hasFunc(key)
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
	index := a.hasFunc(key)
	pairs := a.buckets[index]
	for i, pair := range pairs {
		if pair.key == key {
			a.size--
			a.buckets[index] = append(pairs[:i], pairs[i+1:]...)
		}
	}
}

func (a *ArrayHashMap) hasFunc(key int) int {
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
		loadThres:      3.0 / 5.0,
		extendRatio:    2,
	}
}

//go:generate go run main.go
func main() {
	hashMap := NewArrayHashMap(5)
	hashMap.Set(0, "00")
	hashMap.Set(1, "01")
	hashMap.Set(101, "101")
	hashMap.Set(201, "201")
	hashMap.Set(2, "02")
	hashMap.Print()
	hashMap.Del(1)
	hashMap.Set(101, "101_UPDATE")
	hashMap.Print()
	fmt.Println(hashMap.Keys())
	fmt.Println(hashMap.Values())
	fmt.Println(hashMap.Cap())
}
