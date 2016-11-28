package ticker

import (
	"time"
)

// Ticker represents ticker
type Ticker struct {
	Timer  *time.Timer
	Params Parameters
}

// Returns the duration to wait for next tick
func getNextTickDuration(params Parameters) time.Duration {

	now := time.Now()

	// Generating next tick time
	// based on the values passed in parameters
	nextTick := time.Date(now.Year(),
		now.Month(),
		now.Day(),
		params.Hour,
		params.Minute,
		params.Second,
		params.NanoSecond,
		time.Local)

	// If the tick's due time is before current time,
	// add Interval to it to be re-called in [Interval] hours
	if nextTick.Before(now) {
		nextTick = nextTick.Add(params.Interval)
	}

	// Get duration between the current time and the next tick
	duration := nextTick.Sub(now)

	return duration
}

// New creates new ticker
func New(params Parameters) Ticker {

	var ticker Ticker

	ticker.Params = params

	duration := getNextTickDuration(params)

	ticker.Timer = time.NewTimer(duration)

	return ticker
}

// Update updates ticker
func (jt Ticker) Update() {
	jt.Timer.Reset(getNextTickDuration(jt.Params))
}
