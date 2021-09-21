package YLogger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger interface {
	SetLogOutput(output string)
	SetLogLevel(level string)
	T(args ...interface{})
	D(args ...interface{})
	I(args ...interface{})
	W(args ...interface{})
	E(args ...interface{})
	F(args ...interface{})
}

// stdLogger implements the logger interface using the log package.
// There is no need to specify a date/time prefix since stdout and stderr
// are logged in StackDriver with those values already present.
type YLogger struct {
	logLevel  int8
	logOutput int8
}

func (l *YLogger) SetLogOutput(output string) {
	switch strings.ToUpper(output) {
	case "STDERR":
		l.logOutput = 0
	case "STDOUT":
		l.logOutput = 1
	default:
		panic("BAD logOutput " + output)
	}
}

func (l *YLogger) SetLogLevel(level string) {
	switch strings.ToUpper(level) {
	case "TRACE":
		l.logLevel = 0
	case "DEBUG":
		l.logLevel = 1
	case "INFO":
		l.logLevel = 2
	case "WARN":
		l.logLevel = 3
	case "ERROR":
		l.logLevel = 4
	case "FATAL":
		l.logLevel = 5
	default:
		panic("BAD loglevel " + level)
	}
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
	if l.logLevel <= level {
		t := time.Now().Format("2006/1/2 15:04:05")
		lv := l.levelToString(level)

		msg := t + " " + lv + " " + fmt.Sprint(args...)
		switch l.logOutput {
		case 0:
			fmt.Fprintln(os.Stderr, msg)
		case 1:
			fmt.Fprintln(os.Stdout, msg)
		}

	}
}
