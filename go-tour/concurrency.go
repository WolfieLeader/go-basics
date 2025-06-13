package main

import (
	"fmt"
	"time"
)

func BackgroundSay(s string) {
	for i := range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("BG-%d: %s\n", i, s)
	}
}

func Say(s string) {
	for i := range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("N-%d: %s\n", i, s)
	}
}

func GoRoutinesExample() {
	go BackgroundSay("Foo")
	Say("Bar")
}
