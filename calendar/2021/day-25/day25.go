package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
)

func main() {
	input := files.ReadFile(25, 2021, "\n")
	println(solvePart1(input))
}

func solvePart1(input []string) int {
		e, s, d := parseInput(input)

		didMove, i := true, 0
		for didMove {
			e, s, didMove = step(e, s, d)
			i++
		}

    return i
}

func step(eastHerd sets.Set, southHerd sets.Set, dimensions grid.Coords) (sets.Set, sets.Set, bool) {
	newEast, newSouth := eastHerd.Copy(), southHerd.Copy()
	didMove := false

	eastIterator := eastHerd.Iterator()
	for _, c := range eastIterator {
		coords := grid.ParseCoords(c)
		newCoords := grid.Coords{ X: (coords.X + 1) % dimensions.X, Y: coords.Y}.ToString()
		if !eastHerd.Has(newCoords) && !southHerd.Has(newCoords) {
			newEast.Remove(c)
			newEast.Add(newCoords)
			didMove = true
		}
	}

	southIterator := southHerd.Iterator()
	for _, c := range southIterator {
		coords := grid.ParseCoords(c)
		newCoords := grid.Coords{ X: coords.X, Y: (coords.Y + 1) % dimensions.Y}.ToString()
		if !newEast.Has(newCoords) && !southHerd.Has(newCoords) {
			newSouth.Remove(c)
			newSouth.Add(newCoords)
			didMove = true
		}
	}

	return newEast, newSouth, didMove
}

var east, south = '>', 'v'
func parseInput(input []string) (sets.Set, sets.Set, grid.Coords) {
	eastHerd, southHerd := sets.New(), sets.New()

	for r, row := range input {
		for c, value := range row {
			if value == east {
				eastHerd.Add(grid.Coords{ X: c, Y: r }.ToString())
			} else if value == south {
				southHerd.Add(grid.Coords{ X: c, Y: r }.ToString())
			}
		}
	}

	return eastHerd, southHerd, grid.Coords{ X: len(input[0]), Y: len(input)}
}
