package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"fmt"
	"math"
)

func main() {
	input := files.ReadFile(23, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	elves, _ := simulate(parseInput(input), 10)
	return getBoundingBoxAreas(elves) - elves.Size()
}

func solvePart2(input []string) int {
	_, rounds := simulate(parseInput(input), 999999)
	return rounds + 1
}

func simulate(elves sets.Set, maxRounds int) (sets.Set, int) {
	round := 0
	for round < maxRounds {
		iter := elves.Iterator()
		proposed, proposedCount := map[string]string{}, map[string]int{}
		someMoving := false
		for _, elf := range iter {
			p := propose(elf, elves, round)
			proposed[elf] = p.ToString()
			proposedCount[p.ToString()]++
			if !someMoving && p.ToString() != elf {
				someMoving = true
			}
		}
		if !someMoving {
			return elves, round
		}
		oldSize := elves.Size()
		elves = sets.New()
		for old, new := range proposed {
			if proposedCount[new] >= 2 {
				elves.Add(old)
			} else {
				elves.Add(new)
			}
		}
		if oldSize != elves.Size() {
			panic("we lost some elves")
		}
		round++
	}
	return elves, round
}


var toCheck [4][3]grid.Coords = [4][3]grid.Coords{
	{{ X: -1, Y: -1 }, { X: 0, Y: -1 }, { X: 1, Y: -1 }},
	{{ X: -1, Y: 1 }, { X: 0, Y: 1 }, { X: 1, Y: 1 }},
	{{ X: -1, Y: 0 }, { X: -1, Y: -1 }, { X: -1, Y: 1 }},
	{{ X: 1, Y: 0 }, { X: 1, Y: -1 }, { X: 1, Y: 1 }},
}
var newMove [4]grid.Coords = [4]grid.Coords{
	{ X: 0, Y: -1 },
	{ X: 0, Y: 1 },
	{ X: -1, Y: 0 },
	{ X: 1, Y: 0 },
}
func propose(elf string, occupied sets.Set, round int) grid.Coords {
	elfCoords := grid.ParseCoords(elf)

	if !hasAnyNeighbors(elf, occupied) {
		return elfCoords
	}
	
	starting := round % len(toCheck)
	for i := 0; i < len(toCheck); i++ {
		index := (i + starting) % len(toCheck)
		checking := toCheck[index]
		validMove := true
		for _, c := range checking {
			checked := grid.Coords{ X: elfCoords.X + c.X, Y: elfCoords.Y + c.Y }
			if occupied.Has(checked.ToString()) {
				validMove = false
			}
		}

		if validMove {
			return grid.Coords{ X: elfCoords.X + newMove[index].X, Y: elfCoords.Y + newMove[index].Y } 
		}
	}
	
	return elfCoords
}

func hasAnyNeighbors(elf string, occupied sets.Set) bool {
	elfCoords := grid.ParseCoords(elf)
	for _, list := range toCheck {
		for _, c := range list {
			checked := grid.Coords{ X: elfCoords.X + c.X, Y: elfCoords.Y + c.Y }
			if occupied.Has(checked.ToString()) {
				return true
			}
		}
	}
	return false
}

func parseInput(input []string) sets.Set {
	elves := sets.New()
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				elves.Add(fmt.Sprintf("%d,%d", x, y))
			}
		}
	}
	return elves
}

func getBoundingBoxAreas(elves sets.Set) int {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	iter := elves.Iterator()
	for _, elf := range iter {
		c := grid.ParseCoords(elf)
		minX = maths.Min(minX, c.X)
		maxX = maths.Max(maxX, c.X)
		minY = maths.Min(minY, c.Y)
		maxY = maths.Max(maxY, c.Y)
	}
	return ((maxX - minX) + 1) * ((maxY - minY) + 1)
}

func printElves(elves sets.Set) {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	iter := elves.Iterator()
	for _, elf := range iter {
		c := grid.ParseCoords(elf)
		minX = maths.Min(minX, c.X)
		maxX = maths.Max(maxX, c.X)
		minY = maths.Min(minY, c.Y)
		maxY = maths.Max(maxY, c.Y)
	}
	for y := minY; y <= maxY; y++ {
		line := ""
		for x := minX; x <= maxX; x++ {
			if elves.Has(fmt.Sprintf("%d,%d", x, y)) {
				line += "#"
			} else {
				line += "."
			}
		}
		println(line)
	}
}
