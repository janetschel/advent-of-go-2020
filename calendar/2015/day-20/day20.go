package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"math"
	"strconv"
)

func main() {
	input := files.ReadFile(20, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	target, _ := strconv.Atoi(input[0])
	for houseNumber := 1; houseNumber <= target; houseNumber++ {
		if maths.SumOfDivisors(houseNumber)*10 >= target {
			return houseNumber
		}
	}

	return 0
}

func solvePart2(input []string) int {
	target, _ := strconv.Atoi(input[0])
	for houseNumber := 1; houseNumber <= target; houseNumber++ {
		min := int(math.Floor(float64(houseNumber / 50)))
		divisors := maths.Divisors(houseNumber)
		sum := 0
		for i := range divisors {
			if divisors[i] > min {
				sum += divisors[i]
			}
		}
		if sum*11 >= target {
			return houseNumber
		}
	}

	return 0
}
