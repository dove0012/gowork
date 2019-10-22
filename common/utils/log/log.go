package log

import (
	"gowork/common/utils/time"
	l "log"
	"runtime"
)

//Info log info
func Info(msg string, v ...interface{}) {
	str := "log info: " + msg + "\n"
	if v != nil {
		l.Printf(str, v)
	} else {
		l.Printf(str)
	}

}

//Warn log warn
func Warn(msg string, v ...interface{}) {
	str := "log warning: " + msg + "\n"
	if v != nil {
		l.Printf(str, v)
	} else {
		l.Printf(str)
	}
}

//Debug log debug
func Debug(msg string, v ...interface{}) {
	str := "log debug: " + msg + "\n"
	if v != nil {
		l.Printf(str, v)
	} else {
		l.Printf(str)
	}
}

//Error log error
func Error(msg string, v ...interface{}) {
	str := "log Error: " + msg + "\n"
	if v != nil {
		l.Printf(str, v)
	} else {
		l.Printf(str)
	}
}

//Error2Exit log error2Exit
func Error2Exit(msg string, v ...interface{}) {
	str := "log Error2Exit: " + msg + "\n"
	if v != nil {
		l.Printf(str, v)
	} else {
		l.Printf(str)
	}
	runtime.Goexit()
}

//TimeConsuming TimeConsuming
func TimeConsuming(starTime int64, msg string) {
	l.Printf("%s TimeConsuming: %f second\n", msg, float32(time.NowUnixMilli()-starTime)/1000)
}
