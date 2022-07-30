package gosql

import "time"

func timeStr() string {
	Time := time.Now().Format("2006-01-02 15:04:05")
	return Time
}
