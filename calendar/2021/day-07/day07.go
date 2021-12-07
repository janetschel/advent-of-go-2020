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

	return calcTotalMoves(numbers, median)
}

func solvePart2(input []string) int {
	splits := strings.Split(input[0], ",")
	numbers := []int{}
	for _, val := range splits {
		numVal, _ := strconv.Atoi(val)
		numbers = append(numbers, numVal)
	}
	floor, ceil := calcMean(numbers)
	ceilVal := calcTotalMovesWithIncrease(numbers, ceil)
	floorVal := calcTotalMovesWithIncrease(numbers, floor)
	return int(math.Min(float64(ceilVal), float64(floorVal)))
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

func calcMean(n []int) (int, int) {
	total := 0

	for _, v := range n {
		total += v
	}

	// IMPORTANT: return was rounded!
	return int(math.Floor(float64(total) / float64(len(n)))), int(math.Ceil(float64(total) / float64(len(n))))

}
