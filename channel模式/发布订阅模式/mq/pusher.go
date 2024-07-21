package mq

import (
	"fmt"
	"sync"
	"time"
)

type MQ interface {
	SubScribe(topic MatchTopic) Subscriber
	Push(message *Message)
	Close()
	Delete(subscriber Subscriber)
}

type Pusher struct {
	//消息存储长度
	messageSize int

	//订阅关系和主题匹配
	subscribers map[Subscriber]MatchTopic

	//超时时间
	timeOut time.Duration

	//读写锁
	m sync.RWMutex
}

func NewPusher(size int, timeOut time.Duration) MQ {
	return &Pusher{
		messageSize: size,
		subscribers: make(map[Subscriber]MatchTopic),
		timeOut:     timeOut,
		m:           sync.RWMutex{},
	}
}

func (p *Pusher) SubScribe(topic MatchTopic) Subscriber {
	subscriber := make(Subscriber, p.messageSize)
	p.m.Lock()
	defer p.m.Unlock()
	p.subscribers[subscriber] = topic
	return subscriber
}

func (p *Pusher) Push(message *Message) {
	p.m.Lock()
	defer p.m.Unlock()
	var wg sync.WaitGroup

	for subscriber, topic := range p.subscribers {
		wg.Add(1)
		go p.dispatchMessage(subscriber, message, topic, &wg)
	}

	wg.Wait()
}

func (p *Pusher) dispatchMessage(subscriber Subscriber, message *Message, topic MatchTopic, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(message) {
		return
	}
	select {
	case subscriber <- message:
	case <-time.After(p.timeOut):
		fmt.Println("发送超时！！！")
	}
}

func (p *Pusher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for subscriber, _ := range p.subscribers {
		delete(p.subscribers, subscriber)
		close(subscriber)
	}
}

func (p *Pusher) Delete(subscriber Subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, subscriber)
	close(subscriber)
}
