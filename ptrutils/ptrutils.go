package ptrutils

import "time"

func Pstring(str string) *string {
	return &str
}

func Pbool(boolean bool) *bool {
	return &boolean
}

func Pint(integer int64) *int64 {
	return &integer
}

func Pfloat(float float64) *float64 {
	return &float
}

func PTimeDuration(timeDuration time.Duration) *time.Duration {
	return &timeDuration
}

func PTime(time time.Time) *time.Time {
	return &time
}
