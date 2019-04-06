package main

import (
	"fmt"

	"github.com/ajpasquale/timer"
)

func main() {
	tm := timer.NewTimer(550)
	for {
		if ticks, ok := tm.Tick(); ok {
			fmt.Println("New unit spawned?")
			if ticks > 4 {
				tm.Stop()
				break
			}
		}
	}
}
