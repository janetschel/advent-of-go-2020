package main

import (
	"advent-of-go/utils/files"
	"strconv"
)

type grid [][]int

type coords struct {
	x int
	y int
}

func main() {
	input := files.ReadFile(11, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	g := parseInput(input)

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += simulateStep(g)
	}

	return flashes
}

func solvePart2(input []string) int {
	g := parseInput(input)

	i := 0
	// arbitrary stop condition
	for !allFlashed(g) && i < 5000 {
		simulateStep(g)
		i++
	}

	return i
}

func parseInput(input []string) grid {
	g := [][]int{}
	for i := range input {
		g = append(g, []int{})
		for j := 0; j < len(input[i]); j++ {
			energy, _ := strconv.Atoi(input[i][j : j+1])
			g[i] = append(g[i], energy)
		}
	}
	return g
}

func simulateStep(g grid) int {
	increment(g)
	flashCoords := getAllFlashCoords(g)
	flashes := 0
	for len(flashCoords) > 0 {
		for _, f := range flashCoords {
			flash(f, g)
			flashes++
		}
		flashCoords = getAllFlashCoords(g)
	}
	expendEnergy(g)
	return flashes
}

func increment(g grid) {
	for y := range g {
		for x := range g[y] {
			g[y][x]++
		}
	}
}

func expendEnergy(g grid) {
	for y := range g {
		for x := range g[y] {
			if g[y][x] == -1 {
				g[y][x] = 0
			}
		}
	}
}

func flash(c coords, g grid) {
	neighbors := getNeighbors(c, g)
	g[c.y][c.x] = -1
	for _, n := range neighbors {
		if g[n.y][n.x] != -1 {
			g[n.y][n.x]++
		}
	}
}

func getAllFlashCoords(g grid) []coords {
	c := []coords{}
	for y := range g {
		for x := range g[y] {
			if g[y][x] > 9 {
				c = append(c, coords{x: x, y: y})
			}
		}
	}
	return c
}

func getNeighbors(c coords, g grid) []coords {
	neighbors := []coords{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			x, y := dx+c.x, dy+c.y
			if x >= 0 && y >= 0 && !(x == c.x && y == c.y) && y < len(g) && x < len(g[y]) {
				neighbors = append(neighbors, coords{x: x, y: y})
			}
		}
	}
	return neighbors
}

func allFlashed(g grid) bool {
	for y := range g {
		for x := range g[y] {
			if g[y][x] != 0 {
				return false
			}
		}
	}
	return true
}
