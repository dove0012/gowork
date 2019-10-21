package log

import (
	"gowork/common/utils/time"
	l "log"
	"runtime"
)

func Error2Exit(err error, msg string) {
	if err != nil {
		l.Printf("%s: %s\n", msg, err)
		runtime.Goexit()
	}
}

func Info(msg string) {
	l.Printf("log info: %s\n", msg)
}

func Warn(msg string) {
	l.Printf("log warning: %s\n", msg)
}

func Debug(msg string) {
	l.Printf("log debug: %s\n", msg)
}

func Error(err error, msg string) {
	if err != nil {
		l.Printf("%s: %s\n", msg, err)
	}
}

func TimeConsuming(starTime int64, msg string) {
	l.Printf("%s TimeConsuming: %f second\n", msg, float32(time.NowUnixMilli() - starTime) / 1000)
}
