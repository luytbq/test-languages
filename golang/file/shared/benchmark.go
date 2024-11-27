package shared

import "time"

type FFunc func()

func ElapsedTime(f FFunc) time.Duration {
	startTime := time.Now()

	f()

	return time.Since(startTime)
}
