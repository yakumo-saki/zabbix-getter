package YLogger

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	T(args ...interface{})
	D(args ...interface{})
	I(args ...interface{})
	W(args ...interface{})
	E(args ...interface{})
	F(args ...interface{})
}

type YLogger struct {
	name      string
	logOutput *int8
	logLevel  *int8
}

func (l *YLogger) T(args ...interface{}) {
	l.log(0, args...)
}

func (l *YLogger) D(args ...interface{}) {
	l.log(1, args...)
}

func (l *YLogger) I(args ...interface{}) {
	l.log(2, args...)
}

func (l *YLogger) W(args ...interface{}) {
	l.log(3, args...)
}

func (l *YLogger) E(args ...interface{}) {
	l.log(4, args...)
}

func (l *YLogger) F(args ...interface{}) {
	l.log(5, args...)
}

func (l *YLogger) levelToString(level int8) string {
	switch level {
	case 0:
		return "TRACE"
	case 1:
		return "DEBUG"
	case 2:
		return "INFO "
	case 3:
		return "WARN "
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	default:
		panic("BAD loglevel " + fmt.Sprint(level))
	}
}

// internal log output implementation
func (l *YLogger) log(level int8, args ...interface{}) {
	if *l.logLevel <= level {
		t := time.Now().Format("2006/1/2 15:04:05")
		lv := l.levelToString(level)

		msg := t + " " + l.name + " " + lv + " " + fmt.Sprint(args...)
		switch *l.logOutput {
		case 0:
			fmt.Fprintln(os.Stderr, msg)
		case 1:
			fmt.Fprintln(os.Stdout, msg)
		}

	}
}
