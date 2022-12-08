package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// A - Rock    - X - 1
// B - Paper   - Y - 2
// C - Scissor - Z - 3

// TIE - 3
// WIN - 6
// Lose - 0

func main() {

	content, err := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n") // Used to Split Lines from the Input FIle

	score := 0 // default score - 0
	for i := 0; i < len(lines); i++ {
		score += scoreForARound(lines[i]) // running sum of all the scores for each line
	}
	fmt.Println("Total Score is ", score)
	if err != nil {
		log.Fatal(err)
	}
}

func scoreForARound(move string) int {
	scores := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7} // Combinations of how the string can occur, one look up
	score := scores[move]
	return score // returning score for one roundf
}
