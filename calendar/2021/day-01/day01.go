package main

import (
	"strconv"
	"tblue-aoc-2021/utils/files"
)

func main() {
	input := files.ReadFile(01, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	numberInput := make([]int, len(input))
	for i, v := range input {
		intVar, _ := strconv.Atoi(v)
		numberInput[i] = intVar
	}
	numIncreases := 0
	for i := 1; i < len(numberInput); i++ {
		if numberInput[i-1] < numberInput[i] {
			numIncreases++
		}
	}

	return numIncreases
}

func solvePart2(input []string) int {
	numberInput := make([]int, len(input))
	for i, v := range input {
		intVar, _ := strconv.Atoi(v)
		numberInput[i] = intVar
	}
	numIncreases := 0
	for i := 2; i < len(numberInput)-1; i++ {
		firstSum := numberInput[i-2] + numberInput[i-1] + numberInput[i]
		secondSum := numberInput[i-1] + numberInput[i] + numberInput[i+1]
		if firstSum < secondSum {
			numIncreases++
		}
	}

	return numIncreases
}
