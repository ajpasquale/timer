package timer

import (
	"testing"
)

func TestNewTimer(t *testing.T) {
	timer := NewTimer(1)
	if timer == nil {
		t.Fatal("timer should not be nil")
	}
}

func TestTick(t *testing.T) {
	timer := NewTimer(0)
	if i, ok := timer.Tick(); ok {
		if i != timer.ticks {
			t.Fatal("tick method should return ticks")
		}
	} else {
		t.Fatal("timer should have ticked")
	}
}

func TestStop(t *testing.T) {
	timer := NewTimer(0)
	timer.Stop()
	if _, ok := timer.Tick(); ok {
		t.Fatal("timer should not have ticked")
	}
}

func TestStart(t *testing.T) {
	timer := NewTimer(0)
	timer.Stop()
	timer.Start()
	if _, ok := timer.Tick(); !ok {
		t.Fatal("timer should have ticked")
	}
}
