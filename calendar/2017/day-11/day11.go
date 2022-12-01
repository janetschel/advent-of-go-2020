package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"strings"
)

type direction string
const (
	north direction = "n"
	northeast = "ne"
	southeast = "se"
	south = "s"
	southwest = "sw"
	northwest = "nw"
)

type axialCoords struct {
	q int
	r int
}

func main() {
	input := files.ReadFile(11, 2017, "\n")
	part1, part2 := solve(input)
	println(part1)
	println(part2)
}

func solve(input []string) (int, int) {
	coordsDeltaMap := map[direction]axialCoords{
		north: { q: 0, r: -1 } ,
		northeast: { q: 1, r: -1 },
		southeast: { q: 1, r: 0 },
		south: { q: 0, r: 1 },
		southwest: { q: -1, r: 1 },
		northwest: { q: -1, r: 0 },
	}

	
	current, origin := axialCoords{q: 0, r: 0}, axialCoords{q: 0, r: 0}
	maxDistance := 0
	distance := 0
	directions := parseInput(input)
	for _, d := range directions {
		delta := coordsDeltaMap[d]
		current.q += delta.q
		current.r += delta.r
		distance = axialDistance(origin, current)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return distance, maxDistance
}

func parseInput(input []string) []direction {
	parts := strings.Split(input[0], ",")
	directions := make([]direction, len(parts))
	for i, d := range parts {
		directions[i] = direction(d)
	}
	return directions
}

func axialDistance(from axialCoords, to axialCoords) int {
	return (int(maths.Abs(from.q - to.q)) +
		int(maths.Abs(from.q + from.r - to.q - to.r)) +
		int(maths.Abs(from.r - to.r))) / 2
}
