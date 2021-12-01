package main

import (
	"advent-of-go/utils/files"
)

type slope struct {
	x int
	y int
}

func main() {
	input := files.ReadFile(03, 2020, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func traverseForest(forest []string, slope slope) int {
	trees := 0

		for y, x := 0, 0; y < len(forest); y, x = y + slope.y, x + slope.x {
			x = x % len(forest[y])
			if (forest[y][x:x+1] == "#") {
				trees++
			}
		}

    return trees
}

func solvePart1(input []string) int {
	return traverseForest(input, slope { x: 3, y: 1 })
}

func solvePart2(input []string) int {
	slopes := make([]slope, 0)
	slopes = append(slopes, slope { x: 1, y: 1 })	
	slopes = append(slopes, slope { x: 3, y: 1})
	slopes = append(slopes, slope { x: 5, y: 1 })
	slopes = append(slopes, slope { x: 7, y: 1 })
	slopes = append(slopes, slope { x: 1, y: 2 })

	result := 1

	for _, s := range slopes {
		result *= traverseForest(input, s)
	}

	return result
}
