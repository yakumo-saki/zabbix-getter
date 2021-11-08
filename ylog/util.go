package ylog

import "fmt"

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
