package main

//TODO 完成协程池

type TaskHandler func(p interface{})

// Task 定义 task
type Task struct {
	Params  interface{}
	Handler TaskHandler
}

// WorkerPool 协程池结构
type WorkerPool struct {
	//队列缓冲大小
	s int
	//task 队列
	q chan Task
}
