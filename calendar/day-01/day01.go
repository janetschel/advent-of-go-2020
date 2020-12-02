package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"errors"
	"fmt"
	"log"
)

func main() {
	inputSliceAsString := files.ReadFile(1, "\n")
	input := conv.ToIntSlice(inputSliceAsString)

	solution, err := findPairToMakeSum(input, 2020)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}

func findPairToMakeSum(input []int, sum int) (int, error) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] + input[j] == sum {
				return input[i] * input[j], nil
			}
		}
	}

	return 0, errors.New(fmt.Sprintf("No pair found to make sum of %d", sum))
}
