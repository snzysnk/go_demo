package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type Queue struct {
	list     *list.List
	capacity int
	m        sync.Mutex
}

var (
	QueueFullErr  = errors.New("队列已满")
	QueueEmptyErr = errors.New("队列已空")
)

func NewQueue(size int) *Queue {
	return &Queue{list: list.New(), capacity: size}
}

func (q *Queue) Push(v interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()
	if q.list.Len() > q.capacity {
		return QueueFullErr
	}
	q.list.PushBack(v)
	return nil
}

func (q *Queue) Pop() (interface{}, error) {
	q.m.Lock()
	defer q.m.Unlock()
	if q.list.Len() <= 0 {
		return nil, QueueEmptyErr
	}
	front := q.list.Front()
	q.list.Remove(front)
	return front.Value, nil
}

func (q *Queue) Dump() {
	fmt.Println(q.list)
}

//go:generate go run main.go
func main() {
	q := NewQueue(3)
	pop, err := q.Pop()
	if err != nil {
		fmt.Println(err) //队列已空
	}
	fmt.Println(pop) // nil

	_ = q.Push(1)
	_ = q.Push(2)
	_ = q.Push(3)

	i, err := q.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	i, err = q.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
