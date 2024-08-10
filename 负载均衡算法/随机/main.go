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
		Address: []string{"IP_ADDRESS:8080", "IP_ADDRESS:8081", "IP_ADDRESS:8082"},
	}
}

func (b *Balancer) Get() string {
	b.m.RLock()
	defer b.m.RUnlock()
	return b.Address[rand.Intn(len(b.Address))]
}

func main() {
	b := NewBalance()
	for i := 0; i < 10; i++ {
		fmt.Println(b.Get())
	}
}
