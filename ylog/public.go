package ylog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Logging struct {
	logLevel   int8
	logOutput  int8
	outputType int8
}

const TIME_FORMAT = "2006/01/02 15:04:05"

var logging Logging

// ファイル名の共通部分（/home/.../OfuroNotifyGo/ のような部分）
var projectBaseDir string

// Must call this from file in project root dir
func Init() {
	_, filepath, _, ok := runtime.Caller(1)
	if !ok {
		projectBaseDir = ""
	}
	fn := path.Base(filepath)
	projectBaseDir = strings.Replace(filepath, fn, "", 1)
}

func GetLogger() *YLogger {
	_, filepath, _, _ := runtime.Caller(1)
	// fn := runtime.FuncForPC(pointer)
	// if !ok || fn == nil {
	// 	return GetLoggerByName("UNKNOWN")
	// }

	// if false {
	// 	return GetLoggerByName(fn.Name())
	// }

	fp := strings.Replace(filepath, projectBaseDir, "", 1)
	return GetLoggerByName(fp)

}

// Loggerを取得します
// name = 名称。 main とか sub1 とか。 ログに出力される。
func GetLoggerByName(name string) *YLogger {
	logger := &YLogger{
		name:      name,
		logOutput: &logging.logOutput,
		logLevel:  &logging.logLevel,
	}

	return logger
}

// ログ出力先を設定します
// output = [STDERR | STDOUT]
func SetLogOutput(output string) {
	switch strings.ToUpper(output) {
	case LOG_OUTPUT_STDERR:
		logging.logOutput = 0
	case LOG_OUTPUT_STDOUT:
		logging.logOutput = 1
	default:
		panic("BAD logOutput " + output)
	}
}

// ログ出力しきい値を設定します。
// level = [TRACE | DEBUG | INFO | WARN | ERROR | FATAL]
func SetLogLevel(level string) error {
	switch strings.ToUpper(level) {
	case LOG_LEVEL_TRACE:
		logging.logLevel = 0
	case LOG_LEVEL_DEBUG:
		logging.logLevel = 1
	case LOG_LEVEL_INFO:
		logging.logLevel = 2
	case LOG_LEVEL_WARN:
		logging.logLevel = 3
	case LOG_LEVEL_ERROR:
		logging.logLevel = 4
	case LOG_LEVEL_FATAL:
		logging.logLevel = 5
	default:
		return fmt.Errorf("SetLogLevel: BAD loglevel %s", level)
	}
	return nil
}

// ログ出力形式を設定します。
// logtype = [PLAIN | JSON]
func SetLogType(logtype string) {
	switch strings.ToUpper(logtype) {
	case LOG_TYPE_PLAIN:
		logging.outputType = OUTPUT_PLAIN
	case LOG_TYPE_JSON:
		logging.outputType = OUTPUT_JSON
	case LOG_TYPE_LTSV:
		logging.outputType = OUTPUT_LTSV
	default:
		panic("SetLogType: BAD logtype " + logtype)
	}
}
