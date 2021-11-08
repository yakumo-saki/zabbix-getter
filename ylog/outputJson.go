package ylog

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/imdario/mergo"
)

type awsJsonLog struct {
	Time    string                 `json:"time"`
	Level   string                 `json:"level"`
	Name    string                 `json:"name"`
	Message string                 `json:"message"`
	Values  map[string]interface{} `json:"values"`
}

// internal log output implementation
func (l *YLogger) logOutputJson(level int8, args ...interface{}) {
	if *l.logLevel <= level {
		t := time.Now().Format(TIME_FORMAT)
		lv := l.levelToString(level)

		var v map[string]interface{}
		mergo.Merge(&v, withValues, mergo.WithOverride)
		mergo.Merge(&v, l.values, mergo.WithOverride)

		out := awsJsonLog{
			Time:   t,
			Level:  strings.Trim(lv, " "),
			Name:   l.name,
			Values: v,
		}

		if len(args) > 0 {
			out.Message = fmt.Sprint(args...)
		}

		bytes, _ := json.Marshal(out)
		msg := string(bytes)

		switch *l.logOutput {
		case 0:
			fmt.Fprintln(os.Stderr, msg)
		case 1:
			fmt.Fprintln(os.Stdout, msg)
		}
	}
}
