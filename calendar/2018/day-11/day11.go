package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := files.ReadFile(11, 2018, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) string {
	serialNumber, _ := strconv.Atoi(input[0])
	powerLevels := make([][]int, 300)

	for y := range powerLevels {
		powerLevels[y] = make([]int, 300)
		for x := range powerLevels[y] {
			powerLevels[y][x] = calculatePowerLevel(grid.Coords{ X: x + 1, Y: y + 1 }, serialNumber)
		}
	}

	maxCoords, max := grid.Coords{}, math.MinInt
	for y := 0; y < len(powerLevels) - 2; y++ {
		for x := 0; x < len(powerLevels[y]) - 2; x++ {
			sum := getTotalPowerLevel(powerLevels, grid.Coords{ X: x, Y: y }, 3)
			if sum > max {
				max = sum
				maxCoords.X = x + 1
				maxCoords.Y = y + 1
			}
		}
	}

	return maxCoords.ToString()
}

// pretty sure this should be a DP solution but the brute force works and I didn't want to research
func solvePart2(input []string) string {
	serialNumber, _ := strconv.Atoi(input[0])
	powerLevels := make([][]int, 300)

	for y := range powerLevels {
		powerLevels[y] = make([]int, 300)
		for x := range powerLevels[y] {
			powerLevels[y][x] = calculatePowerLevel(grid.Coords{ X: x + 1, Y: y + 1 }, serialNumber)
		}
	}

	maxCoords, maxSize, max := grid.Coords{}, 0, math.MinInt
	for size := 1; size <= len(powerLevels); size++ {
		for y := 0; y < len(powerLevels) - (size - 1); y++ {
			for x := 0; x < len(powerLevels[y]) - (size - 1); x++ {
				sum := getTotalPowerLevel(powerLevels, grid.Coords{ X: x, Y: y }, size)
				if sum > max {
					max = sum
					maxCoords.X = x + 1
					maxCoords.Y = y + 1
					maxSize = size
				}
			}
		}
	}

	return fmt.Sprintf("%v,%d", maxCoords.ToString(), maxSize)
}

func calculatePowerLevel(c grid.Coords, serialNumber int) int {
	rackID := c.X + 10
	powerLevel := rackID * c.Y
	powerLevel += serialNumber
	powerLevel *= rackID
	powerLevel = getHundredsDigit(powerLevel)
	return powerLevel - 5
}

func getHundredsDigit(number int) int {
	return number  / 100 % 10
}

func getTotalPowerLevel(powerLevels [][]int, topLeft grid.Coords, size int) int {
	sum := 0
	for dy := 0; dy < size; dy++ {
		for dx := 0; dx < size; dx++ {
			sum += powerLevels[topLeft.Y+dy][topLeft.X+dx]
		}
	}
	return sum
}
