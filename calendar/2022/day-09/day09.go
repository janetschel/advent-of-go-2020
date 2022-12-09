package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"strconv"
	"strings"
)

var right, left, up, down = "R", "L", "U", "D"
type move struct {
	direction string
	steps int
}
var directionMap = map[string]grid.Coords { right: { X: 1, Y: 0}, left: { X: -1, Y: 0 }, up: { X: 0, Y: 1 }, down: { X: 0, Y: -1 }}
var differenceMap = map[string]grid.Coords {
	"0,0": { X: 0, Y: 0},
	"0,1": { X: 0, Y: 0},
	"0,-1": { X: 0, Y: 0},
	"1,0": { X: 0, Y: 0},
	"-1,0": { X: 0, Y: 0},
	"1,1": { X: 0, Y: 0},
	"1,-1": { X: 0, Y: 0},
	"-1,1": { X: 0, Y: 0},
	"-1,-1": { X: 0, Y: 0},
	"0,2": { X: 0, Y: 1},
	"0,-2": { X: 0, Y: -1},
	"2,0": { X: 1, Y: 0},
	"-2,0": { X: -1, Y: 0},
	"2,1": { X: 1, Y: 1},
	"2,-1": { X: 1, Y: -1},
	"-2,1": { X: -1, Y: 1},
	"-2,-1": { X: -1, Y: -1},
	"1,2": { X: 1, Y: 1},
	"1,-2": { X: 1, Y: -1},
	"-1,2": { X: -1, Y: 1},
	"-1,-2": { X: -1, Y: -1},
	"2,2": { X: 1, Y: 1},
	"2,-2": { X: 1, Y: -1},
	"-2,-2": { X: -1, Y: -1},
	"-2,2": { X: -1, Y: 1},
}

func main() {
	input := files.ReadFile(9, 2022, "\n")
	println(moveKnots(input, 2))
	println(moveKnots(input, 10))
}

func moveKnots(input []string, knotCount int) int {
	tailPositions := sets.New()
	moves := parseInput(input)
	knots := make([]grid.Coords, knotCount)
	for i := range knots {
		knots[i] = grid.Coords{ X: 0, Y: 0 }
	}

	tailPos := len(knots) - 1
	tailPositions.Add(knots[tailPos].ToString())
	for _, m := range moves {
		v := directionMap[m.direction]
		for i := 0; i < m.steps; i++ {
			knots[0].X += v.X
			knots[0].Y += v.Y
			for j := 1; j < len(knots); j++ {
				key := grid.Coords{ X: knots[j-1].X - knots[j].X, Y: knots[j-1].Y - knots[j].Y }.ToString()
				v1, exists := differenceMap[key]
				if !exists {
					panic("unexpected motion " + key)
				}
				knots[j].X += v1.X
				knots[j].Y += v1.Y
				if j == tailPos {
					tailPositions.Add(knots[j].ToString())
				}
			}
		}
	}

	return tailPositions.Size()
}

func parseInput(input []string) []move {
	moves := make([]move, len(input))

	for i, line := range input {
		parts := strings.Fields(line)
		steps, _ := strconv.Atoi(parts[1])
		moves[i] = move{ direction: parts[0], steps: steps }
	}

	return moves
}
