package main

import (
	"advent-of-go/utils/files"
	"strconv"
)

func main() {
	input := files.ReadFile(01, 2017, "\n")

	digits := make([]int, len(input[0]))
	for i, d := range input[0] {
		digit, _ := strconv.Atoi(string(d))
		digits[i] = digit
	}

	println(inverseCapcha(digits, 1))
	println(inverseCapcha(digits, len(digits) / 2))
}

func inverseCapcha(input []int, indexModifier int) int {
	total := 0
	for i, curr := range input {
		if curr == elementAtCircular(input, i+indexModifier) {
			total += curr
		}
	}

	return total
}

func elementAtCircular(arr []int, index int) int {
	return(arr[index % len(arr)])
}
