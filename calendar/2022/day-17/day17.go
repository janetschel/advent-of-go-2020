package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"fmt"
	"strings"
)

func main() {
	input := files.ReadFile(17, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	highestPoint, _ := simulate(input[0], 2022)
	return highestPoint
}

func solvePart2(input []string) int {
	highestPoint, _ := simulate(input[0], 10000)
	return highestPoint
}

func simulate(wind string, rocks int) (int, []int) {
	delta, highestPoint := make([]int, rocks), 0

	w := 0
	occupied := sets.New()
	for round := 0; round < rocks; round++ {
		shape, movingDown := getShape(round, highestPoint), true
		for movingDown {
			if wind[w] == '>' {
				shape, _ = moveRight(shape, occupied)
			} else {
				shape, _ = moveLeft(shape, occupied)
			}
			w = (w + 1) % len(wind)
			shape, movingDown = moveDown(shape, occupied)
		}
		for _, p := range shape {
			if p.Y > highestPoint {
				highestPoint = p.Y
			}
			occupied.Add(p.ToString())
		}
		delta[round] = highestPoint - slices.Sum(delta)
	}

	fmt.Println(delta)

	return highestPoint, delta
}

func getShape(round int, highestPoint int) []grid.Coords {
	highestPoint = highestPoint + 1
	switch round % 5 {
	case 0:
		// --
		return []grid.Coords{
			{ X: 2, Y: highestPoint + 3 },
			{ X: 3, Y: highestPoint + 3 },
			{ X: 4, Y: highestPoint + 3 },
			{ X: 5, Y: highestPoint + 3 },
		}
	case 1:
		// +
		return []grid.Coords{
			{ X: 3, Y: highestPoint + 3 },
			{ X: 2, Y: highestPoint + 4 },
			{ X: 3, Y: highestPoint + 4 },
			{ X: 4, Y: highestPoint + 4 },
			{ X: 3, Y: highestPoint + 5 },
		}
	case 2:
		// _|
		return []grid.Coords{
			{ X: 2, Y: highestPoint + 3 },
			{ X: 3, Y: highestPoint + 3 },
			{ X: 4, Y: highestPoint + 3 },
			{ X: 4, Y: highestPoint + 4 },
			{ X: 4, Y: highestPoint + 5 },
		}
	case 3:
		// |
		return []grid.Coords{
			{ X: 2, Y: highestPoint + 3 },
			{ X: 2, Y: highestPoint + 4 },
			{ X: 2, Y: highestPoint + 5 },
			{ X: 2, Y: highestPoint + 6 },
		}
	case 4:
		// █
		return []grid.Coords{
			{ X: 2, Y: highestPoint + 3 },
			{ X: 3, Y: highestPoint + 3 },
			{ X: 2, Y: highestPoint + 4 },
			{ X: 3, Y: highestPoint + 4 },
		}
	}
	panic("bad round")
}

func moveRight(shape []grid.Coords, occupied sets.Set) ([]grid.Coords, bool) {
	newShape := make([]grid.Coords, len(shape))

	for i, c := range shape {
		newC := grid.Coords{ X: c.X + 1, Y: c.Y }
		if newC.X > 6 || occupied.Has(newC.ToString()) {
			shapeCopy := make([]grid.Coords, len(shape))
			copy(shapeCopy, shape)
			return shapeCopy, false
		}
		newShape[i] = newC
	}

	return newShape, true
}

func moveLeft(shape []grid.Coords, occupied sets.Set) ([]grid.Coords, bool) {
	newShape := make([]grid.Coords, len(shape))

	for i, c := range shape {
		newC := grid.Coords{ X: c.X - 1, Y: c.Y }
		if newC.X < 0 || occupied.Has(newC.ToString()) {
			shapeCopy := make([]grid.Coords, len(shape))
			copy(shapeCopy, shape)
			return shapeCopy, false
		}
		newShape[i] = newC
	}

	return newShape, true
}

func moveDown(shape []grid.Coords, occupied sets.Set) ([]grid.Coords, bool) {
	newShape := make([]grid.Coords, len(shape))

	for i, c := range shape {
		newC := grid.Coords{ X: c.X, Y: c.Y - 1 }
		if newC.Y <= 0 || occupied.Has(newC.ToString()) {
			shapeCopy := make([]grid.Coords, len(shape))
			copy(shapeCopy, shape)
			return shapeCopy, false
		}
		newShape[i] = newC
	}

	return newShape, true
}

func printCave(occupied sets.Set, highestPoint int) {
	for y := highestPoint; y > 0; y-- {
		line := fmt.Sprintf("%2d |", y)
		for x := 0; x < 7; x++ {
			if occupied.Has(fmt.Sprintf("%d,%d", x, y)) {
				line += "#"
			} else {
				line += "."
			}
		}
		line += "|"
		println(line)
		if !strings.Contains(line, "#") {
			panic("bad cave")
		}
	}
	println("   +-------+")
	println()
}
