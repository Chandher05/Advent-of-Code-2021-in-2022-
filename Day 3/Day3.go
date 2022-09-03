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
	sum := make([]int, 12)

	for i := 0; i < len(lines); i++ {
		oneRow := strings.Split(lines[i], "")
		// fmt.Println(oneRow)

		for j := 0; j < len(oneRow); j++ {
			if parseInt(oneRow[j]) == 1 {
				sum[j] += 1
			}
		}
	}
	// fmt.Println("total values", sum)
	// fmt.Println("max value", len(lines))
	gamma := ""
	epsilon := ""

	for k := 0; k < len(sum); k++ {
		if sum[k] > (len(lines) / 2) {
			gamma += "1"
			epsilon += "0"
		} else {
			epsilon += "1"
			gamma += "0"
		}
		// fmt.Println(gamma, epsilon, k, sum[k])
	}
	fmt.Println(gamma, epsilon)

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println("Answer", gammaInt*epsilonInt)

}

func parseInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}

// 011100101100
// 100011010011
