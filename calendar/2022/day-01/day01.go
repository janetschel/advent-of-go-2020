package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"sort"
	"strconv"
)

func main() {
	input := files.ReadFile(1, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	calories := calculateCalories(input)
	return calories[len(calories) - 1]
}

func solvePart2(input []string) int {
	calories := calculateCalories(input)

	return slices.Sum(calories[len(calories) - 3:])
}

func calculateCalories(input []string) []int {
	calories, currentSum := []int{}, 0
	for _, value := range input {
		if value == "" {
			calories = append(calories, currentSum)
			currentSum = 0
		} else {
			calories, _ := strconv.Atoi(value)
			currentSum += calories
		}
	}

	sort.Ints(calories)
	return calories
}
