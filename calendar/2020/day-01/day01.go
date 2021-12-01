package main

import (
	"advent-of-go/utils/files"
	"log"
	"strconv"
)

func main() {
	input := files.ReadFile(01, 2020, "\n")
	numbers := parseToIntSlice(input);
	println(solvePart1(numbers))
	println(solvePart2(numbers))
}

func parseToIntSlice(input []string) []int {
	numbers := make([]int, len(input))
	for i, str := range input {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = num
	}
	return numbers
}

func solvePart1(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == 2020 {
				return numbers[i] * numbers[j]
			}
		}
	}

	return 0
}

func solvePart2(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i] + numbers[j] + numbers[k] == 2020 {
					return numbers[i] * numbers[j] * numbers[k]
				}
			}
		}
	}

	return 0
}
