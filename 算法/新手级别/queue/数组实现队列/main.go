package main

import (
	"errors"
	"fmt"
	"sync"
)

type Queue struct {
	front    int
	size     int
	capacity int
	data     []interface{}
	m        sync.Mutex
}

var (
	QueueFullErr  = errors.New("队列已满")
	QueueEmptyErr = errors.New("队列已空")
)

func NewQueue(capacity int) *Queue {
	return &Queue{
		front:    0,
		size:     0,
		capacity: capacity,
		data:     make([]interface{}, capacity),
	}
}

func (q *Queue) isFull() bool {
	return q.size == q.capacity
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

// 入队
func (q *Queue) push(v interface{}) error {
	q.m.Lock()
	defer q.m.Unlock()
	if q.isFull() {
		return QueueFullErr
	}
	//环形队列
	index := (q.front + q.size) % q.capacity
	q.data[index] = v
	q.size++
	return nil
}

// 出队
func (q *Queue) pop() (interface{}, error) {
	q.m.Lock()
	defer q.m.Unlock()
	if q.isEmpty() {
		return nil, QueueEmptyErr
	}
	value := q.data[q.front]
	q.front++
	q.size--
	return value, nil
}

func main() {
	queue := NewQueue(3)
	_, err := queue.pop()
	if err == QueueEmptyErr {
		fmt.Println(err)
	}

	if err = queue.push(1); err != nil {
		fmt.Println(1, err)
	}

	if err = queue.push(2); err != nil {
		fmt.Println(2, err)
	}

	if err = queue.push(3); err != nil {
		fmt.Println(3, err)
	}

	if err = queue.push(4); err != nil {
		fmt.Println(4, err)
	}

	v, _ := queue.pop()
	fmt.Println(v)

	if err = queue.push(4); err != nil {
		fmt.Println(4, err)
	}
}
