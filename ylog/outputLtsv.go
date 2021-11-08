package ylog

import (
	"fmt"
	"os"
	"time"
)

// internal log output implementation
func (l *YLogger) logOutputLtsv(level int8, args ...interface{}) {
	if *l.logLevel <= level {
		t := time.Now().Format(TIME_FORMAT)
		lv := l.levelToString(level)

		msg := fmt.Sprintf("time:%s\tlevel:%s\tname:%s", t, lv, l.name)

		for k, v := range withValues {
			msg = fmt.Sprintf("%s\t%s:%s", msg, k, v)
		}

		if len(args) > 0 {
			msg = fmt.Sprintf("%s\tmessage:%s", msg, fmt.Sprint(args...))
		}

		for k, v := range l.values {
			msg = fmt.Sprintf("%s\t%s:%s", msg, k, v)
		}

		switch *l.logOutput {
		case 0:
			fmt.Fprintln(os.Stderr, msg)
		case 1:
			fmt.Fprintln(os.Stdout, msg)
		}

	}
}
