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

	horizontal := 0
	depth := 0
	aim := 0

	for i := 0; i < len(lines); i++ {
		command := strings.Split(lines[i], " ")
		// fmt.Println(command[0], "by", command[1])

		if command[0] == "forward" {
			horizontal += parseInt(command[1])
			depth += (aim * parseInt(command[1]))
			// fmt.Println(command[0], "by", command[1])

			// fmt.Println("Horizontal", horizontal)
		}
		if command[0] == "up" {
			aim -= parseInt(command[1])

			// fmt.Println("Depth (Up)", depth)
		}
		if command[0] == "down" {
			aim += parseInt(command[1])

			// fmt.Println("Depth (Down)", depth)
		}

	}
	fmt.Println("total Value", depth*horizontal)
}

func parseInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}
