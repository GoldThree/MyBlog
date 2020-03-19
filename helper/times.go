package helper

import "time"

func TimeToTimestamp(t time.Time) int64 {
	return t.UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
}
