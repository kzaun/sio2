package sio2

import "time"

func Unix2Date(uninx int64) string {
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(uninx, 0).Format(timeLayout)
}
