package main

import (
	"advent-of-go-2020/utils"
	"strconv"
)

func main() {
	input := utils.ReadFile(1, "\n")

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				first, _ := strconv.Atoi(input[i])
				second, _ := strconv.Atoi(input[j])
				third, _ := strconv.Atoi(input[k])

				if first + second + third == 2020 {
					println(first * second * third)
					return
				}
			}
		}
	}
}
