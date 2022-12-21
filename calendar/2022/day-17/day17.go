package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"fmt"
	"strings"
)

func main() {
	input := files.ReadFile(17, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	return simulate(input[0], 2022)
}

func solvePart2(input []string) int {
	return simulate(input[0], 1000000000000)
}

func simulate(wind string, totalRocks int) int {
	highestPoint, extraHeight := 0, 0
	w := 0
	occupied := sets.New()

	cache := map[string][2]int{}
	for rock := 0; rock < totalRocks; rock++ {
		shape, movingDown := getShape(rock, highestPoint), true
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

		// the first ten rows are somewhat arbitrarily selected as part of the cache key
		key := fmt.Sprintf("%v %v", getCacheKey(rock, w), caveToString(occupied, highestPoint, highestPoint - 10))
		results, seen := cache[key]
		if seen {
			dRocks, dHeight := rock - results[1], highestPoint - results[0]
			cyclePeriod := ((totalRocks - results[1]) / dRocks) - 1
			extraHeight += cyclePeriod * dHeight
			rock += cyclePeriod * dRocks
		} else {
			cache[key] = [2]int{ highestPoint, rock }
		}
	}

	return highestPoint + extraHeight
}

func getCacheKey(round int, w int) string {
	return fmt.Sprintf("%d,%d", round % 5, w)
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

func caveToString(occupied sets.Set, highestPoint int, floor int) string {
	cave := ""
	for y := highestPoint; y > maths.Max(floor, 0); y-- {
		line := "|"
		for x := 0; x < 7; x++ {
			if occupied.Has(fmt.Sprintf("%d,%d", x, y)) {
				line += "#"
			} else {
				line += "."
			}
		}
		line += "|"
		cave += line + "\n"
		if !strings.Contains(line, "#") {
			panic("bad cave")
		}
	}
	cave += "+-------+" + "\n"
	return cave
}
