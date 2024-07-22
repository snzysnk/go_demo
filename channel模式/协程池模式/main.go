package main

import (
	"fmt"
	"sync"
	"time"
)

//TODO 完成协程池

type TaskHandler func(p interface{})

// Task 定义 task
type Task struct {
	Params  interface{}
	Handler TaskHandler
}

type WorkerPoolAbility interface {
	AddWorker()
	Release()
	AddTask(task Task)
}

// WorkerPool 协程池结构
type WorkerPool struct {
	ch chan Task
	wg sync.WaitGroup
}

func (w *WorkerPool) AddTask(task Task) {
	w.ch <- task
}

func (w *WorkerPool) AddWorker() {
	w.wg.Add(1)
	go func() {
		for t := range w.ch {
			time.Sleep(3 * time.Second)
			t.Handler(t.Params)
		}
		w.wg.Done()
	}()
}

func (w *WorkerPool) Release() {
	close(w.ch)
	w.wg.Wait()
}

func main() {
	pool := WorkerPool{
		ch: make(chan Task, 100),
		wg: sync.WaitGroup{},
	}
	pool.AddWorker()
	pool.AddWorker()
	pool.AddTask(Task{
		Params: "Hello",
		Handler: func(p interface{}) {
			fmt.Println(p)
		},
	})
	pool.AddTask(Task{
		Params: "World!",
		Handler: func(p interface{}) {
			fmt.Println(p)
		},
	})
	pool.AddTask(Task{
		Params: "满了!!!",
		Handler: func(p interface{}) {
			fmt.Println(p)
		},
	})
	pool.Release()
}
