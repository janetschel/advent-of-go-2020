package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"errors"
)

func main() {
	input := files.ReadFile(9, "\n")

	// target from part 1
	result, err := solvePart2(conv.ToIntSlice(input), 776203571)

	if err != nil {
		panic(err)
	}

	println("Result of min+max of range that sum to target:", result)
}

func solvePart2(input []int, target int) (int, error) {

	for i := 0; i < len(input); i++ {
		min := input[i]
		max := input[i]
		sum := input[i]

		for j := i + 1; j < len(input); j++ {
			sum += input[j]
			if sum == target {
				return min + max, nil
			}

			if input[j] > max {
				max = input[j]
			} else if input[j] < min {
				min = input[j]
			}
		}
	}

    return 0, errors.New("no number-range in whole input that sum to target")
}
