package xlog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	LUTC           = 1 << iota //使用UTC时间
	Lprefix                    //前缀放在最前面
	Ldate                      //需要日期
	Ltime                      //需要时分秒
	Lmicrosecnonds             //需要微秒时间
	Lfile                      //记录文件名和位置信息
	LDefault       = Ldate | Ltime | Lmicrosecnonds | Lfile | Lprefix
)

type ILog interface {
	SetOutput(io.Writer)
	SetPrefix(string)
	Prefix() string
	Flags() int
	SetFlags(int)
	output(pc uintptr, calldepth int, appendOutput func([]byte) []byte) error
	Output(calldepth int, s string) error
}

type Logger struct {
	outMu     sync.Mutex
	out       io.Writer
	prefix    atomic.Pointer[string]
	flag      atomic.Int32
	isDiscard atomic.Bool
}

func (l *Logger) Output(calldepth int, s string) error {
	calldepth++
	return l.output(0, calldepth, func(b []byte) []byte {
		return append(b, s...)
	})
}

func (l *Logger) output(pc uintptr, calldepth int, appendOutput func([]byte) []byte) error {
	var (
		flag = l.Flags()
		line int
		file string
		ok   bool
	)
	if flag&Lfile != 0 {
		if pc == 0 {
			_, file, line, ok = runtime.Caller(calldepth)
			if !ok {
				file = "???"
				line = 0
			}
		} else {
			fs, _ := runtime.CallersFrames([]uintptr{pc}).Next()
			file = fs.File
			if file == "" {
				file = "???"
			}
			line = fs.Line
		}
	}

	buf := getBuffer()
	defer putBuffer(buf)
	formatHeader(buf, time.Now(), l.Prefix(), flag, file, line)
	*buf = appendOutput(*buf)
	l.outMu.Lock()
	defer l.outMu.Unlock()
	_, err := l.out.Write(*buf)
	return err
}

var bufferPool = sync.Pool{New: func() any { return new([]byte) }}
var std = New(os.Stderr, "^@^", LDefault)

func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func getBuffer() *[]byte {
	buffer := bufferPool.Get().(*[]byte)
	*buffer = (*buffer)[:0]
	return buffer
}

func putBuffer(buffer *[]byte) {
	if cap(*buffer) > 64<<10 {
		*buffer = nil
	}
	bufferPool.Put(buffer)
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
	l.out = writer
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

func formatHeader(buf *[]byte, t time.Time, prefix string, flag int, file string, line int) {
	if flag&Lprefix != 0 {
		*buf = append(*buf, prefix...)
	}

	if flag&LUTC != 0 {
		t = t.UTC()
	}

	if flag&Ldate != 0 {
		year, month, day := t.Date()
		itoa(buf, year, 4)
		*buf = append(*buf, '/')
		itoa(buf, int(month), 2)
		*buf = append(*buf, '/')
		itoa(buf, int(day), 2)
		*buf = append(*buf, ' ')
	}

	if flag&(Ltime|Lmicrosecnonds) != 0 {
		hour, min, sec := t.Clock()
		itoa(buf, hour, 2)
		*buf = append(*buf, ':')
		itoa(buf, min, 2)
		*buf = append(*buf, ':')
		itoa(buf, sec, 2)
		if flag&Lmicrosecnonds != 0 {
			*buf = append(*buf, '.')
			itoa(buf, t.Nanosecond()/1e3, 6)
		}
		*buf = append(*buf, ' ')
	}

	if flag&Lfile != 0 {
		*buf = append(*buf, file...)
		*buf = append(*buf, ": "...)
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}

	if flag&Lprefix == 0 {
		*buf = append(*buf, prefix...)
	}
}
