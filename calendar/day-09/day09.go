package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"errors"
)

func main() {
	input := files.ReadFile(9, "\n")
	result, err := solvePart1(conv.ToIntSlice(input))

	if err != nil {
		panic(err)
	}

	println("First number with no pair of sum in preamble:", result)
}

func solvePart1(input []int) (int, error) {
	preamble := 25

	for i := preamble; i < len(input); i++ {
		valid := false

		for j := i - preamble; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if input[j] + input[k] == input[i] {
					valid = true
				}
			}
		}

		if !valid {
			return input[i], nil
		}
	}

	return 0, errors.New("no number without sum in preamble found")
}
