package main

import (
	"fbdy/mq"
	"fmt"
	"sync"
	"time"
)

func main() {
	pusher := mq.NewPusher(10, 5*time.Second)
	scribeOne := pusher.SubScribe(func(message *mq.Message) bool {
		return message.Topic == "one"
	})
	scribeTwo := pusher.SubScribe(func(message *mq.Message) bool {
		return message.Topic == "two"
	})

	scribeAll := pusher.SubScribe(nil)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for message := range scribeOne {
			if message != nil {
				fmt.Println("one", message.Data)
			} else {
				fmt.Println("one quit")
			}
		}
		wg.Done()
	}()

	go func() {
		for message := range scribeTwo {
			if message != nil {
				fmt.Println("two", message.Data)
			} else {
				fmt.Println("two quit")
			}
		}
		wg.Done()
	}()

	go func() {
		for message := range scribeAll {
			time.Sleep(1 * time.Second)
			if message != nil {
				fmt.Println("all", message.Data)
			} else {
				fmt.Println("all quit")
			}
		}
		wg.Done()
	}()

	pusher.Push(&mq.Message{
		Topic: "one",
		Data:  "hello one",
	})
	pusher.Push(&mq.Message{
		Topic: "two",
		Data:  "hello two",
	})

	pusher.Push(&mq.Message{
		Topic: "two",
		Data:  "hello two",
	})
	pusher.Push(&mq.Message{
		Topic: "two",
		Data:  "hello two",
	})
	pusher.Push(&mq.Message{
		Topic: "two",
		Data:  "hello two",
	})

	pusher.Close()
	wg.Wait()
}
