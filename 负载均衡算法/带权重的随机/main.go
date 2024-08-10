package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Balancer struct {
	Address []string
	m       sync.RWMutex
}

func NewBalance() *Balancer {
	return &Balancer{
		Address: []string{},
	}
}

func (b *Balancer) Get() string {
	b.m.RLock()
	defer b.m.RUnlock()
	return b.Address[rand.Intn(len(b.Address))]
}

func (b *Balancer) Add(address string, weight int) {
	b.m.Lock()
	defer b.m.Unlock()
	for i := 0; i < weight; i++ {
		b.Address = append(b.Address, address)
	}
}

func main() {
	b := NewBalance()
	b.Add("127.0.0.1:9001", 1)
	b.Add("127.0.0.1:9002", 100)
	b.Add("127.0.0.1:9003", 49)
	for i := 0; i < 10; i++ {
		fmt.Println(b.Get())
	}
}
