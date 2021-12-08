package main

import (
	"advent-of-go/utils/files"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func main() {
	input := files.ReadFile(18, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := simulate(input, 100, false)
	return countLightsTurnedOn(result)
}

func solvePart2(input []string) int {
	result := simulate(input, 100, true)
	return countLightsTurnedOn(result)
}

func countLightsTurnedOn(grid []string) int {
	count := 0
	for row := range grid {
		count += strings.Count(grid[row], "#")
	}
	return count
}

func simulate(grid []string, steps int, cornersOn bool) []string {
	gridCopy := make([]string, len(grid))
	copy(gridCopy, grid)
	for i := 0; i < steps; i++ {
		gridCopy = simulateOneStep(gridCopy, cornersOn)
	}
	return gridCopy
}

func simulateOneStep(grid []string, cornersOn bool) []string {
	newGrid := []string{}
	minCoords, maxCoords := Coordinate { x: 0, y: 0}, Coordinate { x: len(grid[0]) - 1, y: len(grid) - 1 }
	for row := range grid {
		newRow := ""
		for column := range grid[row] {
			current := Coordinate{ x: column, y: row }
			if IsCorner(current, grid) {
				newRow += "#"
			} else {
				neighbors := GetNeighbors(current, minCoords, maxCoords)
				neighborsOn := 0
				for i := range neighbors {
					if (IsOn(neighbors[i], grid, cornersOn)) {
						neighborsOn++
					}
				}
				if neighborsOn == 3 ||
					(IsOn(current, grid, cornersOn) && neighborsOn == 2) {
						newRow += "#"
				} else {
					newRow += "."
				}
			}
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func IsCorner(coords Coordinate, grid []string) bool {
	return (coords.x == 0 && coords.y == 0) ||
		(coords.x == 0 && coords.y == len(grid) - 1) ||
		(coords.x == len(grid[0]) - 1 && coords.y == 0) ||
		(coords.x == len(grid[0]) - 1 && coords.y == len(grid) - 1)
}

func IsOn(coords Coordinate, grid []string, cornersOn bool) bool {
	if cornersOn && IsCorner(coords, grid) {
		return true
	}
	light := grid[coords.y][coords.x:coords.x + 1]
	return light == "#"
}

func GetNeighbors(coords Coordinate, minCoords Coordinate, maxCoords Coordinate) []Coordinate {
	neighbors := []Coordinate {}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			newX, newY := x + coords.x, y + coords.y
			if !(x == 0 && y == 0) &&
				newX <= maxCoords.x && newX >= minCoords.x &&
				newY <= maxCoords.y && newY >= minCoords.y {
				neighbors = append(neighbors, Coordinate{ x: newX, y: newY })
			}
		}
	}

	return neighbors
}
