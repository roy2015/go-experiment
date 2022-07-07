package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	tick := time.NewTicker(1 * 1000 * time.Millisecond)
	//tick := time.Tick(1 * 1000 * time.Millisecond)
	go func() {
		for {
			select {
			case c := <-tick.C:
				fmt.Println(c)
			}
		}
	}()

	time.Sleep(10 * 1000 * time.Millisecond)
	tick.Stop()
}
