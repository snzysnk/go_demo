package main

import (
	"fmt"
	"sync"
)

type Balancer struct {
	Address []string
	Next    int
	m       sync.RWMutex
}

func NewBalance() *Balancer {
	return &Balancer{
		Address: []string{"IP_ADDRESS:8080", "IP_ADDRESS:8081", "IP_ADDRESS:8082"},
		Next:    0,
	}
}

func (b *Balancer) Get() string {
	b.m.RLock()
	defer b.m.RUnlock()
	addr := b.Address[b.Next]
	b.Next = (b.Next + 1) % len(b.Address)
	return addr
}

func main() {
	balancer := NewBalance()
	fmt.Println(balancer.Get()) //IP_ADDRESS:8080
	fmt.Println(balancer.Get()) //IP_ADDRESS:8081
	fmt.Println(balancer.Get()) //IP_ADDRESS:8082
	fmt.Println(balancer.Get()) //IP_ADDRESS:8080
}
