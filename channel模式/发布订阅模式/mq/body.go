package mq

type Message struct {
	Topic string `json:"topic"`
	Data  interface{}
}
