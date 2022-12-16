package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(15, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	pairs, minX, maxX, maxDist := parseInput(input)
	invalid := findInvalidLocationsInRow(2000000, pairs, minX, maxX, maxDist)
	return invalid.Size()
}

func findInvalidLocationsInRow(row int, pairs [][2]grid.Coords, minX int, maxX int, maxDist int) sets.Set {
	invalid := sets.New() 
	for x := minX - maxDist; x <= maxX + maxDist; x++ {
		c := grid.Coords{ X: x, Y: row }
		checkPairs:
			for _, p := range pairs {
				sensor, beacon := p[0], p[1]
				if c.X == beacon.X && c.Y == beacon.Y {
					break checkPairs
				}
				if sensor.ManhattanDistance(beacon) >= sensor.ManhattanDistance(c) {
					invalid.Add(c.ToString())
					break checkPairs
				}
			}
	}
	return invalid
}

func solvePart2(input []string) int {
	min, max := 0, 4000000

	pairs, _, _, _ := parseInput(input)
	for y := min; y <= max; y++ {
		for x := min; x <= max; x++ {
			couldBeDistress := true
			c := grid.Coords{X: x, Y: y}
			for i := 0; i < len(pairs) && couldBeDistress; i++ {
				sensor, beacon := pairs[i][0], pairs[i][1]
				sensorRange := sensor.ManhattanDistance(beacon)
				if sensorRange >= sensor.ManhattanDistance(c) {
					couldBeDistress = false
					// increase x by sensor "diamond" width
					x += (sensorRange - maths.Abs(sensor.Y - c.Y)) + (sensor.X - c.X)
				}
			}
			if couldBeDistress {
				println(c.ToString())
				return (x * 4000000) + y
			}
		}
	}

	return -1
}

func parseInput(input []string) ([][2]grid.Coords, int, int, int) {
	pairs := make([][2]grid.Coords, len(input))
	minX, maxX, maxDist := math.MaxInt, math.MinInt, math.MinInt
	for i, line := range input {
		fields := strings.FieldsFunc(line, func(char rune) bool {
			return strings.ContainsRune(", =:", char)
		})
		sx, _ := strconv.Atoi(fields[3])
		sy, _ := strconv.Atoi(fields[5])
		bx, _ := strconv.Atoi(fields[11])
		by, _ := strconv.Atoi(fields[13])
		pairs[i] = [2]grid.Coords{{X: sx, Y: sy}, {X: bx, Y: by}}
		if sx < minX {
			minX = sx
		}
		if bx < minX {
			minX = bx
		}
		if sx > maxX {
			maxX = sx
		}
		if bx > maxX {
			maxX = bx
		}
		d := pairs[i][0].ManhattanDistance(pairs[i][1]) 
		if d > maxDist {
			maxDist = d
		}
	}
	return pairs, minX, maxX, maxDist
}
