package timer

import (
	"fmt"
	"testing"
)

func TestNewTimer(t *testing.T) {
	tr := NewTimer(5, 550, true)
	for {
		if ok := tr.Tick(); ok {
			fmt.Println("tick")
			break
		}
	}
}
