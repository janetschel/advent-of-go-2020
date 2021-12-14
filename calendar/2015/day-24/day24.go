package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"math"
	"strconv"
)

func main() {
	input := files.ReadFile(24, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	return findMinimumGroup(parseInput(input), 3)
}

func solvePart2(input []string) int {
	return findMinimumGroup(parseInput(input), 4)
}

func parseInput(input []string) []int {
	nums := []int{}
	for i := range input {
		num, _ := strconv.Atoi(input[i])
		nums = append(nums, num)
	}
	return nums
}

func quantumEntanglement(input []int) int {
	qe := 1
	for i := range input {
		qe *= input[i]
	}
	return qe
}

func findMinimumGroup(input []int, n int) int {
	target := slices.Sum(input) / n
	minLength := 0

	validGroups := [][]int{}
	for i := 0; i < len(input) && minLength == 0; i++ {
		combinations := slices.GenerateCombinationsLengthN(input, i)
		for _, current := range combinations {
			if slices.Sum(current) == target {
				minLength = len(current)
				validGroups = append(validGroups, current)
			}
		}
	}

	minQuantumEntanglement := math.MaxInt
	for i := range validGroups {
		qe := quantumEntanglement(validGroups[i])
		if qe < minQuantumEntanglement {
			minQuantumEntanglement = qe
		}
	}

	return minQuantumEntanglement
}
