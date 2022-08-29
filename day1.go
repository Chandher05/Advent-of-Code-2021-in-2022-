package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("input.txt")

	lines := strings.Split(string(content), "\n")

	if err != nil {
		log.Fatal(err)
	}

	count := 0
	prev := toString(lines[0]) + toString(lines[1]) + toString(lines[2])
	curr := toString(lines[1]) + toString(lines[2]) + toString(lines[3])

	for i := 2; i < len(lines)-2; i++ {
		if curr > prev {
			count = count + 1
		}
		println(curr, prev, count)

		prev = curr
		curr = toString(lines[i]) + toString(lines[i+1]) + toString(lines[i+2])

	}
	if curr > prev {
		count = count + 1
	}
	fmt.Println(count)

}

func toString(s string) int {
	val, _ := strconv.Atoi(s)

	return val
}
