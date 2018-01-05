package time

import (
	t "time"
)

func NowFormat() string {
	return t.Now().Format("2006-01-02 15:04:05")
}

//返回当前时间秒
func NowUnix() int64 {
	return t.Now().Unix()
}

//返回当前时间毫秒
func NowUnixMilli() int64 {
	return t.Now().UnixNano() / 1e6
}

//返回当前时间纳秒
func NowUnixNano() int64 {
	return t.Now().UnixNano()
}
