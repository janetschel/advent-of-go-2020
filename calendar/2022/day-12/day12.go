package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"sort"
)

func main() {
	input := files.ReadFile(12, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	puzzle, start, end := parseInput(input)
	return findShortestPath(puzzle, []grid.Coords{ start }, end)
}

func solvePart2(input []string) int {
	puzzle, _, end := parseInput(input)
	starts := []grid.Coords{}
	for y, line := range puzzle {
		for x, char := range line {
			if char == 'a' {
				starts = append(starts, grid.Coords{ X: x, Y: y })
			}
		}
	}
	return findShortestPath(puzzle, starts, end)
}

func findShortestPath(puzzle [][]rune, starts []grid.Coords, end grid.Coords) int {
	shortestPaths := map[string]*sets.Set{}
	queue := sets.New()
	for _, start := range starts {
		startingSet := sets.New()
		startingSet.Add(start.ToString())
		shortestPaths[start.ToString()] = &startingSet
		queue.Add(start.ToString())
	}
	winningPathLengths := []int{}

	killswitch := 0
	for queue.Size() > 0 && killswitch < 99999999 {
		currentStr := queue.Random()
		queue.Remove(currentStr)
		current := grid.ParseCoords(currentStr)

		if current.X == end.X && current.Y == end.Y {
			path := shortestPaths[currentStr]
			winningPathLengths = append(winningPathLengths, path.Size())
		} else {
			neighbors := []grid.Coords{
				{ X: current.X + 1, Y: current.Y },
				{ X: current.X - 1, Y: current.Y },
				{ X: current.X, Y: current.Y + 1 },
				{ X: current.X, Y: current.Y - 1 },
			}
			visited := shortestPaths[currentStr]
			currentValue := puzzle[current.Y][current.X]
			for _, n := range neighbors {
				newStr := n.ToString()
				shortest, hasVisited := shortestPaths[newStr]
				if grid.IsInGrid(n, puzzle) &&
					currentValue - puzzle[n.Y][n.X] >= -1 &&
					!visited.Has(newStr) &&
					(!hasVisited || visited.Size() + 1 < shortest.Size()) {
							queue.Add(newStr)
							newShortest := visited.Copy()
							newShortest.Add(newStr)
							shortestPaths[newStr] = &newShortest
				}
			}
		}
		killswitch++
	}

	if len(winningPathLengths) == 0 {
		return - 1
	}
	winningPathLengths = slices.Filter(winningPathLengths, func(a int) bool { return a > 0})
	sort.Ints(winningPathLengths)
	return winningPathLengths[0] - 1
}

func parseInput(input []string) ([][]rune, grid.Coords, grid.Coords) {
	var start, end grid.Coords
	puzzle := make([][]rune, len(input))

	for y, line := range input {
		puzzle[y] = make([]rune, len(line))
		for x, char := range line {
			if char == 'S' {
				start = grid.Coords { X: x, Y: y }
				puzzle[y][x] = 'a'
			} else if char == 'E' {
				end = grid.Coords { X: x, Y: y }
				puzzle[y][x] = 'z'
			} else {
				puzzle[y][x] = char
			}
		}
	}

	return puzzle, start, end
}
