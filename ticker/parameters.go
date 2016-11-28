package ticker

import (
	"fmt"
	"time"
)

// Parameters represents the parameters needed
// to make the Ticker work
type Parameters struct {
	Hour       int
	Minute     int
	Second     int
	NanoSecond int
	Interval   time.Duration
}

// Check if ticker parameters are correct
func (p Parameters) Check() error {

	if p.Hour >= 24 || p.Minute >= 60 || p.Second >= 60 {
		return fmt.Errorf("Ticker parameters not good")
	}
	return nil
}
