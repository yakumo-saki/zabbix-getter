package ylog

import "fmt"

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

	values map[string]interface{}
}

const OUTPUT_PLAIN = 0
const OUTPUT_JSON = 1
const OUTPUT_LTSV = 2

const LOG_OUTPUT_STDERR = "STDERR"
const LOG_OUTPUT_STDOUT = "STDOUT"

const LOG_LEVEL_TRACE = "TRACE"
const LOG_LEVEL_DEBUG = "DEBUG"
const LOG_LEVEL_INFO = "INFO"
const LOG_LEVEL_WARN = "WARN"
const LOG_LEVEL_ERROR = "ERROR"
const LOG_LEVEL_FATAL = "FATAL"

const LOG_TYPE_JSON = "JSON"
const LOG_TYPE_PLAIN = "PLAIN"
const LOG_TYPE_LTSV = "LTSV"

var withValues map[string]interface{}

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

func (l *YLogger) log(level int8, args ...interface{}) {
	switch logging.outputType {
	case OUTPUT_PLAIN:
		l.logOutputPlain(level, args...)
	case OUTPUT_JSON:
		l.logOutputJson(level, args...)
	case OUTPUT_LTSV:
		l.logOutputLtsv(level, args...)
	default:
		panic("UNKNOWN OUTPUT : " + fmt.Sprint(logging.outputType))
	}

	if l.values != nil && len(l.values) > 0 {
		l.values = make(map[string]interface{})
	}

}

// json時次のログ出力に追加する値。これでセットした値は一度ログ出力を行うと削除される
func (l *YLogger) Add(key string, value interface{}) *YLogger {
	if l.values == nil {
		l.values = make(map[string]interface{})
	}
	l.values[key] = value

	return l
}

// json時次のログ出力に追加する値。addとの違いはこちらは一度追加したら残り続ける
func (l *YLogger) With(key string, value interface{}) *YLogger {
	if withValues == nil {
		withValues = make(map[string]interface{})
	}
	withValues[key] = value
	return l
}
