package main

import (
	"advent-of-go-2020/utils"
	"strconv"
)

func main() {
	input := utils.ReadFile(1, "\n")

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			first, _ := strconv.Atoi(input[i])
			second, _ := strconv.Atoi(input[j])

			if first + second == 2020 {
				println(first * second)
				return
			}
		}
	}
}

