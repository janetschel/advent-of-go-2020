package main

import (
	"advent-of-go-2020/utils"
)

func main() {
	inputSliceAsString := utils.ReadFile(1, "\n")
	input := utils.ToIntSlice(inputSliceAsString)

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i] + input[j] + input[k] == 2020 {
					println(input[i] * input[j] * input[k])
					return
				}
			}
		}
	}
}
