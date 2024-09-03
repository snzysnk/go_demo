package xlog

import (
	"io"
	"sync"
	"sync/atomic"
)

type ILog interface {
	SetOutput(io.Writer)
	SetPrefix(string)
	Prefix() string
	Flags() int
	SetFlags(int)
}

type Logger struct {
	outMu     sync.Mutex
	out       io.Writer
	prefix    atomic.Pointer[string]
	flag      atomic.Int32
	isDiscard atomic.Bool
}

func (l *Logger) Prefix() string {
	if prefix := l.prefix.Load(); prefix != nil {
		return *prefix
	}
	return ""
}

func (l *Logger) Flags() int {
	return int(l.flag.Load())
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.outMu.Lock()
	defer l.outMu.Unlock()
	l.SetOutput(writer)
}

func (l *Logger) SetPrefix(s string) {
	l.prefix.Store(&s)
}

func (l *Logger) SetFlags(i int) {
	l.flag.Store(int32(i))
}

func New(out io.Writer, prefix string, flag int) *Logger {
	l := new(Logger)
	l.SetOutput(out)
	l.SetPrefix(prefix)
	l.SetFlags(flag)
	return l
}
