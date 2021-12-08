package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"math"
	"strconv"
)

func main() {
	input := files.ReadFile(07, 2021, ",")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	distances := parseDistances(input)
	return findMinimumFuel(distances, calculateFuelLinear)
}

func solvePart2(input []string) int {
	distances := parseDistances(input)
	return findMinimumFuel(distances, calculateFuelSummation)
}

func parseDistances(input []string) []int {
	distances := []int{}
	for i := range input {
		d, _ := strconv.Atoi(input[i])
		distances = append(distances, d)
	}
	return distances
}

func calculateFuelLinear(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

func calculateFuelSummation(a int, b int) int {
	steps := calculateFuelLinear(a, b)
	fuel := 0
	for i := 1; i <= steps; i++ {
		fuel += i
	}
	return fuel
}

func findMinimumFuel(distances []int, calculateFuel func(int, int) int) int {
	minTotal := math.MaxInt
	maxDistance := slices.Max(distances)
	minDistance := slices.Min(distances)
	for i := minDistance; i <= maxDistance; i++ {
		total := 0
		for j := range distances {
			total += calculateFuel(i, distances[j])
		}
		if total < minTotal {
			minTotal = total
		}
	}

	return minTotal
}
