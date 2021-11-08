package ylog

import (
	"fmt"
	"os"
	"time"
)

// internal log output implementation
func (l *YLogger) logOutputPlain(level int8, args ...interface{}) {
	if *l.logLevel <= level {
		t := time.Now().Format(TIME_FORMAT)
		lv := l.levelToString(level)

		msg := t + " " + l.name + " " + lv

		for k, v := range withValues {
			msg = fmt.Sprintf("%s %s=%s", msg, k, v)
		}

		if len(args) > 0 {
			msg = msg + " " + fmt.Sprint(args...)
		}

		for k, v := range l.values {
			msg = fmt.Sprintf("%s %s=%s", msg, k, v)
		}

		switch *l.logOutput {
		case 0:
			fmt.Fprintln(os.Stderr, msg)
		case 1:
			fmt.Fprintln(os.Stdout, msg)
		}

	}
}
