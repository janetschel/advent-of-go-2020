package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(02, 2021, "\n")

	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	x, y := 0, 0

	for i := range input {
		parts := strings.Split(input[i], " ")
		val, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			x += val
		case "down":
			y += val
		case "up":
			y -= val
		}
	}

	fmt.Printf("Position: %v, %v\n", x, y)
	return x * y
}

func solvePart2(input []string) int {
	x, y, aim := 0, 0, 0

	for i := range input {
		parts := strings.Split(input[i], " ")
		val, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			x += val
			y += (aim * val)
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	fmt.Printf("Position: %v, %v, Aim: %v\n", x, y, aim)
	return x * y
}
