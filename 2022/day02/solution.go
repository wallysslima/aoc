package main

import (
	"bufio"
	"fmt"
	"os"
)

// A|X - Rock | B|Y - Paper | C|Z - Scissors
func main() {

	scores := map[string]int{
		// lost
		"B X": 0 + 1,
		"C Y": 0 + 2,
		"A Z": 0 + 3,
		// draw
		"A X": 3 + 1,
		"B Y": 3 + 2,
		"C Z": 3 + 3,
		// win
		"C X": 6 + 1,
		"A Y": 6 + 2,
		"B Z": 6 + 3,
	}
	score := 0

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		score += scores[input.Text()]
	}

	fmt.Println(score)
}
