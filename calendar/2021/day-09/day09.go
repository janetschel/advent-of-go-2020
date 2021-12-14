package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"sort"
	"strconv"
)

type point struct {
	x     int
	y     int
	value int
}

func main() {
	input := files.ReadFile(9, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	risk := 0
	points := parseInput(input)
	lowPoints := findLowPoints(points)
	for i := range lowPoints {
		risk += lowPoints[i].value + 1
	}

	return risk
}

func solvePart2(input []string) int {
	points := parseInput(input)
	lowPoints := findLowPoints(points)
	basinSizes := make([]int, len(lowPoints))
	for i := range lowPoints {
		basinSizes[i] = calculateBasinSize(lowPoints[i], points)
	}
	return multiply(basinSizes)
}

func parseInput(input []string) [][]int {
	ints := [][]int{}
	for i := range input {
		ints = append(ints, make([]int, len(input[i])))
		for j := range input[i] {
			val, _ := strconv.Atoi(input[i][j : j+1])
			ints[i][j] = val
		}
	}
	return ints
}

func findLowPoints(points [][]int) []point {
	lowPoints := []point{}
	for y := range points {
		for x := range points[y] {
			current := points[y][x]
			if current != 9 {
				isLowest := true
				for _, d := range getVisitorDeltas() {
					dx, dy := x+d[0], y+d[1]
					if dx >= 0 && dy >= 0 &&
						dx < len(points[y]) && dy < len(points) &&
						points[dy][dx] < current {
						isLowest = isLowest && points[dy][dx] > current
					}
				}
				if isLowest {
					lowPoints = append(lowPoints, point{x: x, y: y, value: current})
				}
			}
		}
	}
	return lowPoints
}

func calculateBasinSize(lowPoint point, points [][]int) int {
	nextUp := []point{lowPoint}
	visited := []string{}

	size := 0
	for len(nextUp) > 0 {
		current := nextUp[len(nextUp)-1]
		nextUp = nextUp[:len(nextUp)-1]
		if !slices.Contains(visited, current.GetKey()) {
			visited = append(visited, current.GetKey())
			size++

			for _, d := range getVisitorDeltas() {
				dx, dy := current.x+d[1], current.y+d[0]
				if !slices.Contains(visited, fmt.Sprintf("%v,%v", dx, dy)) &&
					dx >= 0 && dy >= 0 &&
					dx < len(points[0]) && dy < len(points) &&
					points[dy][dx] > current.value &&
					points[dy][dx] != 9 {
					nextUp = append(nextUp, point{x: dx, y: dy, value: points[dy][dx]})
				}
			}
		}
	}
	return size
}

/*
Magic Missile costs 53 mana. It instantly does 4 damage.
Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.
*/

func multiply(basins []int) int {
	sort.Ints(basins)
	product := 1
	for i := len(basins) - 1; i >= len(basins)-3; i-- {
		product *= basins[i]
	}
	return product
}

func (p *point) GetKey() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func getVisitorDeltas() [][]int {
	return [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
}
