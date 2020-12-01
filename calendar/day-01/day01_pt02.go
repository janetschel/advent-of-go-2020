package main

import (
	"advent-of-go-2020/utils"
	"errors"
	"fmt"
	"log"
)

func main() {
	inputSliceAsString := utils.ReadFile(1, "\n")
	input := utils.ToIntSlice(inputSliceAsString)

	solution, err := findThreeNumbersToMatchSum(input, 2020)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solution)
}

func findThreeNumbersToMatchSum(input []int, sum int) (int, error) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i] + input[j] + input[k] == sum {
					return input[i] * input[j] * input[k], nil
				}
			}
		}
	}

	return 0, errors.New(fmt.Sprintf("No three numbers found to make sum of %d", sum))
}
