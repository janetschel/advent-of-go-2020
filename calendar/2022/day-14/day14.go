package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"strings"
)

func main() {
	input := files.ReadFile(14, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	rocks, floor := parseInput(input)
	return simulate(rocks, floor, false)
}

func solvePart2(input []string) int {
	rocks, floor := parseInput(input)
	return simulate(rocks, floor + 1, true)
}

func simulate(rocks sets.Set, floor int, hasFloor bool) int {
	blocked := rocks.Copy()
	source := grid.Coords{ X: 500, Y: 0 }
	sand, count := grid.Coords{ X: source.X, Y: source.Y }, 0

	for {
		next := getNextDestination(sand, blocked)
		if next.Y >= floor {
			// abyss
			if !hasFloor {
				return count
			}			
			blocked.Add(next.ToString())
			sand.X = source.X
			sand.Y = source.Y
			count++
		} else if next.X == source.X && next.Y == source.Y {
			return count + 1
		} else if next.X == sand.X && next.Y == sand.Y {
			// at rest
			blocked.Add(next.ToString())
			sand.X = source.X
			sand.Y = source.Y
			count++
		} else {
			// keep moving
			sand.X = next.X
			sand.Y = next.Y
		}
	}
}

func getNextDestination(sand grid.Coords, blocked sets.Set) grid.Coords {
	down, left, right := grid.Coords{ X: sand.X, Y: sand.Y + 1}, grid.Coords{ X: sand.X - 1, Y: sand.Y + 1}, grid.Coords{ X: sand.X + 1, Y: sand.Y + 1 }
	if !blocked.Has(down.ToString()) {
		return down
	}
	
	if !blocked.Has(left.ToString()) {
		return left
	}

	if !blocked.Has(right.ToString()) {
		return right
	}

	return sand
}

func parseInput(input []string) (sets.Set, int) {
	rocks, maxY := sets.New(), 0
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		for i := 1; i < len(parts); i++ {
			c0, c1 := grid.ParseCoords(parts[i-1]), grid.ParseCoords(parts[i])
			var vector grid.Coords
			if c0.X > c1.X {
				vector.X = -1
			} else if c0.X < c1.X {
				vector.X = 1
			} else if c0.Y > c1.Y {
				vector.Y = -1
			} else if c0.Y < c1.Y {
				vector.Y = 1
			}
			c := grid.Coords{ X: c0.X, Y: c0.Y }
			for c.X != c1.X || c.Y != c1.Y {
				if c.Y > maxY {
					maxY = c.Y
				}
				rocks.Add(c.ToString())
				c.X += vector.X
				c.Y += vector.Y
			}
			rocks.Add(c1.ToString())
		}
	}
	return rocks, maxY
}
