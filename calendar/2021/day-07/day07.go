package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

func main() {
	input := files.ReadFile(07, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	splits := strings.Split(input[0], ",")
	numbers := []int{}
	for _, val := range splits {
		numVal, _ := strconv.Atoi(val)
		numbers = append(numbers, numVal)
	}
	median := calcMedian(numbers)
	println(median)

	return calcTotalMoves(numbers, median)
}

func solvePart2(input []string) int {
	splits := strings.Split(input[0], ",")
	numbers := []int{}
	for _, val := range splits {
		numVal, _ := strconv.Atoi(val)
		numbers = append(numbers, numVal)
	}
	mean := calcMean(numbers)
	println(mean)

	return calcTotalMovesWithIncrease(numbers, mean)
}

func calcTotalMovesWithIncrease(positions []int, destination int) int {
	sum := 0
	for _, val := range positions {
		distance := int(math.Abs(float64(destination - val)))
		sum += (distance * (distance + 1)) / 2
	}
	return sum
}

func calcMedian(n []int) int {
	sort.Ints(n) // sort the numbers

	mNumber := len(n) / 2

	if len(n)%2 == 1 {
		return n[mNumber]
	}

	return (n[mNumber-1] + n[mNumber]) / 2
}

func calcTotalMoves(positions []int, destination int) int {
	sum := 0
	for _, val := range positions {
		sum += int(math.Abs(float64(destination - val)))
	}
	return sum
}

func calcMean(n []int) int {
	total := 0

	for _, v := range n {
		total += v
	}
	println(float64(total) / float64(len(n)))

	// IMPORTANT: return was rounded!
	return int(math.Floor(float64(total) / float64(len(n))))

}
