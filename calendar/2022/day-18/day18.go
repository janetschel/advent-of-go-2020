package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type threeDCoords struct {
	x int
	y int
	z int
}

func main() {
	input := files.ReadFile(18, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	cubes := parseInput(input)
	return getSurfaceArea(cubes)
}

func solvePart2(input []string) int {
	cubes := parseInput(input)
	return getExteriorSurfaceArea(cubes)
}

func parseThreeDCoords(input string) threeDCoords {
	parts := strings.Split(input, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return threeDCoords{ x: x, y: y, z: z }
}

func parseInput(input []string) []threeDCoords {
	coords := []threeDCoords{}
	for _, c := range input {
		coords = append(coords, parseThreeDCoords(c))
	}
	return coords
}

func (c threeDCoords) getNeighbors() [6]threeDCoords {
	return [6]threeDCoords{
		{ x: c.x + 1, y: c.y, z: c.z },
		{ x: c.x - 1, y: c.y, z: c.z },
		{ x: c.x, y: c.y + 1, z: c.z },
		{ x: c.x, y: c.y - 1, z: c.z },
		{ x: c.x, y: c.y, z: c.z + 1 },
		{ x: c.x, y: c.y, z: c.z - 1 },
	}
}

func (c threeDCoords) toString() string {
	return fmt.Sprintf("%d,%d,%d", c.x, c.y, c.z)
}

func getBoundingBox(cubes []threeDCoords) (threeDCoords, threeDCoords) {
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	minZ, maxZ := math.MaxInt, math.MinInt

	for _, c := range cubes {
		minX, maxX = maths.Min(minX, c.x), maths.Max(maxX, c.x)
		minY, maxY = maths.Min(minY, c.y), maths.Max(maxY, c.y)
		minZ, maxZ = maths.Min(minZ, c.z), maths.Max(maxZ, c.z)
	}

	return threeDCoords{ x: minX - 1, y: minY - 1, z: minZ - 1 }, threeDCoords{ x: maxX + 1, y: maxY + 1, z: maxZ + 1 }
}

func getSurfaceArea(cubes []threeDCoords) int {
	cubeSet := sets.New()
	for _, c := range cubes {
		cubeSet.Add(c.toString())
	}
	sides := 6 * len(cubes)
	for _, c := range cubes {
		neighbors := c.getNeighbors()
		for _, n := range neighbors {
			key := n.toString()
			if cubeSet.Has(key) {
				sides--
			}
		}
	}

	return sides
}

func getExteriorSurfaceArea(cubes []threeDCoords) int {
	min, max := getBoundingBox(cubes)
	start := threeDCoords{ x: min.x, y: min.y, z: min.z }
	queue := []threeDCoords{ start }
	cubeSet := sets.New()
	for _, c := range cubes {
		cubeSet.Add(c.toString())
	}
	visited := sets.New()
	visited.Add(start.toString())
	air := 0
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		neighbors := current.getNeighbors()
		for _, next := range neighbors {
			key := next.toString()
			if cubeSet.Has(key) {
				air++
			} else if next.x <= max.x && next.x >= min.x && 
				next.y <= max.y && next.y >= min.y &&
				next.z <= max.z && next.z >= min.z &&
				!visited.Has(key) {
				visited.Add(key)
				queue = append(queue, next)
			}
		}
	}

	return air
}
