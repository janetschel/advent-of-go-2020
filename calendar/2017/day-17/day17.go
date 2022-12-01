package main

import (
	"advent-of-go/utils/files"
	"strconv"
)

func main() {
	input := files.ReadFile(17, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	offset, _ := strconv.Atoi(input[0])
	spinlock, currentIndex := generateSpinlockValues(offset, 2017)
	return spinlock[currentIndex+1]
}

func solvePart2(input []string) int {
	offset, _ := strconv.Atoi(input[0])
	nextIndex, next, currentIndex := 0, 0, 0
	for value := 1; value < 50000000; value++ {
		currentIndex = (currentIndex + offset) % value
		if currentIndex == nextIndex {
			next = value
		}
		if currentIndex < nextIndex {
			nextIndex++
		}
		currentIndex++
	}
	return next
}

func generateSpinlockValues(offset int, iterations int) ([]int, int) {
	spinlock, currentIndex := []int{0}, 0
	for value := 1; value <= iterations; value++ {
		currentIndex = circularIndex(currentIndex, offset, len(spinlock)) + 1
		spinlock = append(spinlock[:currentIndex], append([]int{value}, spinlock[currentIndex:]...)...)
	}

	return spinlock, currentIndex
}

func circularIndex(index int, increment int, length int) int {
	return (index + increment) % length
}
