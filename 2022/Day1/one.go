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
	var calCount []int
	lines := strings.Split(string(content), "\n")
	runningCount := 0
	for i := 0; i < len(lines); i++ {

		if lines[i] == "" {
			calCount = append(calCount, runningCount)
			runningCount = 0
		} else {
			valueInteger, err := strconv.Atoi(lines[i])
			if err != nil {
				panic(err)
			}
			runningCount += valueInteger
		}
	}
	index := 0
	for j := 1; j < len(calCount); j++ {
		if calCount[index] < calCount[j] {
			index = j
		}
	}
	for l := 3; l < len(calCount); l++ {
		if calCount[l] > calCount[0] {
			calCount[2] = calCount[1]
			calCount[1] = calCount[0]
			calCount[0] = calCount[l]
		} else if calCount[l] > calCount[1] && calCount[l] != calCount[1] {
			calCount[2] = calCount[1]
			calCount[1] = calCount[l]
		} else if calCount[l] > calCount[2] {
			calCount[2] = calCount[l]
		}
	}

	fmt.Println(index, calCount[index])
	fmt.Println("Top three", calCount[0]+calCount[1]+calCount[2])
	if err != nil {
		log.Fatal(err)
	}
}
