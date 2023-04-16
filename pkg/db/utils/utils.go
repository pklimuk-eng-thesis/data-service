package db

import "time"

func ParseTime(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}
