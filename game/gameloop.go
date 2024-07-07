package game

import (
	"fmt"
	"time"
)

func GameLoop() {
	// Create an infinitely repeating loop with framework fighters

	round := 1
	for  {
		time.Sleep(1000)
		fmt.Printf("Round %d",round)
		round += 1
	}
}

