package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollAllOnes() (numberOfRolls int) {
	roundCounter := 0
	for {
		roundCounter++
		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1
		die3 := rand.Intn(6) + 1
		die4 := rand.Intn(6) + 1
		if die1 == 1 && die2 == 1 && die3 == 1 && die4 == 1 {
			numberOfRolls = roundCounter
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	NUM_ROWS := 20
	NUM_COLS := 25
	for i := 0; i < NUM_ROWS; i++ {
		for j := 0; j < NUM_COLS; j++ {
			numberOfRolls := rollAllOnes()
			fmt.Print(numberOfRolls, ",")
		}
		fmt.Println()
	}
}
