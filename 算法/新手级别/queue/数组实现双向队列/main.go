package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	QueueFullErr  = errors.New("队列已满")
	QueueEmptyErr = errors.New("队列已空")
)

type Queue struct {
	//为队首提供位置信息
	front int
	//实则为尾部提供位置信息
	size     int
	capacity int
	data     []interface{}
	m        sync.Mutex
}

func newQueue(capacity int) *Queue {
	return &Queue{
		front:    0,
		size:     0,
		capacity: capacity,
		data:     make([]interface{}, capacity),
		m:        sync.Mutex{},
	}
}

func (q *Queue) Index(i int) int {
	return (i + q.capacity) % q.capacity
}

// PushFirst 队首入队
func (q *Queue) PushFirst(v interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()
	if q.size == q.capacity {
		return QueueFullErr
	}
	q.front = q.Index(q.front - 1)
	q.data[q.front] = v
	q.size++
	return nil
}

// PushLast 队尾入队
func (q *Queue) PushLast(v interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()
	if q.size == q.capacity {
		return QueueFullErr
	}
	index := q.Index(q.front + q.size)
	q.data[index] = v
	q.size++
	return nil
}

// PopFirst 队首出队
func (q *Queue) PopFirst() (value interface{}, err error) {
	q.m.Lock()
	defer q.m.Unlock()
	if q.size == 0 {
		return nil, QueueEmptyErr
	}
	value, q.data[q.front] = q.data[q.front], nil
	q.front = q.Index(q.front + 1)
	q.size--
	return value, nil
}

// PopLast 队尾出队
func (q *Queue) PopLast() (value interface{}, err error) {
	q.m.Lock()
	defer q.m.Unlock()
	if q.size == 0 {
		return nil, QueueEmptyErr
	}
	index := q.Index(q.front + q.size - 1)
	value, q.data[index], err = q.data[index], nil, nil
	q.size--
	return value, err
}

func (q *Queue) dump() {
	fmt.Println(q.data)
}

func main() {
	queue := newQueue(5)
	if err := queue.PushFirst(1); err != nil {
		fmt.Println(err)
	}
	if err := queue.PushFirst(2); err != nil {
		fmt.Println(err)
	}
	if err := queue.PushFirst(3); err != nil {
		fmt.Println(err)
	}
	value, err := queue.PopFirst()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
	queue.dump()
	if err := queue.PushLast(5); err != nil {
		fmt.Println(err)
	}
	if err := queue.PushFirst(3); err != nil {
		fmt.Println(err)
	}
	if err := queue.PushLast(4); err != nil {
		fmt.Println(err)
	}

	last, err := queue.PopLast()
	fmt.Println(last)
	if err != nil {
		fmt.Println(err)
	}

	queue.dump()
}
