package util

import "time"

func CurrMillSecond() int64 {
	now := time.Now()
	return now.UnixNano() / int64(time.Millisecond/time.Nanosecond)
}
