package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
)

func main() {
	input := files.ReadFile(17, 2015, "\n")
	inputNums := []int {}
	for i := 0; i < len(input); i++ {
		num, _ := strconv.Atoi(input[i])
		inputNums = append(inputNums, num)
	}
	println(solvePart1(inputNums))
	println(solvePart2(inputNums))
}

func solvePart1(containers []int) int {
	combinations := slices.GenerateAllCombinations(containers)
	count := 0
	for i := 0; i < len(combinations); i++ {
		if slices.Sum(combinations[i]) == 150 {
			count++
		}
	}

	return count
}

func solvePart2(containers []int) int {
	combinations := slices.GenerateAllCombinations(containers)
	containerCountMap := map[int]int {}
	for i := 0; i < len(combinations); i++ {
		if slices.Sum(combinations[i]) == 150 {
			containerCountMap[len(combinations[i])]++
		}
	}

	minContainers := len(containers)
	for key := range containerCountMap {
		if key < minContainers {
			minContainers = key
		}
	}

	return containerCountMap[minContainers]
}
