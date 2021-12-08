package main

import (
	"advent-of-go/utils/files"
	"strconv"
)

func main() {
	input := files.ReadFile(01, 2021, "\n")

	depths := make([]int, len(input))
	for i := range input {
		val, err := strconv.Atoi(input[i])
		if err == nil {
			depths[i] = val
		}
	}
	println(solvePart1(depths))
	println(solvePart2(depths))
}

func solvePart1(input []int) int {
	result := 0

	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			result++
		}
	}

	return result
}

func solvePart2(input []int) int {
	result := 0

	previousSum := 0
	for i := 0; i < len(input)-2; i++ {
		sum := input[i] + input[i+1] + input[i+2]
		if sum > previousSum && previousSum != 0 {
			result++
		}
		previousSum = sum
	}

	return result
}
