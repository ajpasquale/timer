package timer

import "time"

// Timer struct represents the state of a timer
type Timer struct {
	last     time.Time
	interval time.Duration
	ticks    int
	paused   bool
}

// NewTimer
func NewTimer(i int) *Timer {
	interval := (time.Duration(i) * 10) * time.Millisecond
	t := &Timer{
		last:     time.Now(),
		interval: interval,
		ticks:    0,
		paused:   false,
	}
	return t
}

// Tick will increment the timer if it passes the interval other return
// false
func (t *Timer) Tick() (int, bool) {
	now := time.Now()
	if now.Sub(t.last) >= t.interval && !t.paused {
		t.last = now
		t.ticks++
		return t.ticks, true
	}
	return t.ticks, false
}

// Stop sets the internal paused variable to true
func (t *Timer) Stop() {
	t.paused = true
}

// Start sets the internal paused variable to false.
// It does not actually trigger the next tick
func (t *Timer) Start() {
	t.paused = false
}
