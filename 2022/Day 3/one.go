package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Original string
	content, err := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	dict2 := "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Split the string in half
	score := 0
	for i := 0; i < len(lines); i++ {
		halfLen := len(lines[i]) / 2

		part1 := lines[i][:halfLen]
		part2 := lines[i][halfLen:]
		fmt.Println(part1, part2, i)

		for _, char1 := range part1 {
			if strings.ContainsRune(part2, char1) {
				score += strings.Index(dict2, string(char1))
				break
			}
		}
	}
	fmt.Println(score, "Part 1 Total Score")

	scorePart2 := 0
	for j := 0; j < len(lines); j = j + 3 {
		fmt.Println(j, j+1, j+2)
		for _, char1 := range lines[j] {
			if strings.ContainsRune(lines[j+1], char1) && strings.ContainsRune(lines[j+2], char1) {
				scorePart2 += strings.Index(dict2, string(char1))
				break
			}
		}
	}
	fmt.Println(scorePart2, "Part 2 Total Score")

	if err != nil {
		log.Fatal(err)
	}
}
