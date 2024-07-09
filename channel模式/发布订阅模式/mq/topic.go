package mq

type MatchTopic func(message *Message) bool
