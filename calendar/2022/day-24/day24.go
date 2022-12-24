package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
)

func main() {
	input := files.ReadFile(24, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	b, vels, width, height, start, target := parseInput(input)
	blizzardsTable := calculateBlizzardsTable(b, vels, width, height)
	return travelTime(blizzardsTable, start, target, width, height, 0)
}

func solvePart2(input []string) int {
	b, vels, width, height, start, target := parseInput(input)
	blizzardsTable := calculateBlizzardsTable(b, vels, width, height)
	leg1 := travelTime(blizzardsTable, start, target, width, height, 0)
	leg2 := travelTime(blizzardsTable, target, start, width, height, leg1)
	return travelTime(blizzardsTable, start, target, width, height, leg2)
}

func travelTime(blizzardsTable map[int]sets.Set, start grid.Coords, target grid.Coords, width int, height int, startTime int) int {
	lcm := maths.Lcm(width, height)
	time := startTime
	queue := sets.New()
	queue.Add(start.ToString())
	for queue.Size() > 0 {
		time++
		nextQueue := sets.New()
		blizzards := blizzardsTable[time % lcm]
		iter := queue.Iterator()
		for _, pStr := range iter {
			currentPosition := grid.ParseCoords(pStr)
			if currentPosition.X == target.X && currentPosition.Y == target.Y {
				return time
			}
			if !blizzards.Has(currentPosition.ToString()) {
				nextQueue.Add(currentPosition.ToString())
			}
			for _, v := range velocities {
				nextPosition := grid.Coords{ X: currentPosition.X + v.X, Y: currentPosition.Y + v.Y }
				if nextPosition.X == target.X && nextPosition.Y == target.Y {
					return time
				}
				if !blizzards.Has(nextPosition.ToString()) &&
					nextPosition.X >= 0 && nextPosition.X < width && nextPosition.Y >= 0 && nextPosition.Y < height {
					nextQueue.Add(nextPosition.ToString())
				}
			}
			queue = nextQueue
		}
	}
	return time
}

func calculateBlizzardsTable(blizzards []grid.Coords, vels []grid.Coords, width int, height int) map[int]sets.Set {
	table := map[int]sets.Set{}
	lcm := maths.Lcm(width, height)
	for i := 0; i < lcm; i++ {
		s := sets.New()
		for _, k := range blizzards {
			s.Add(k.ToString())
		}
		table[i] = s
		blizzards = moveBlizzards(blizzards, vels, width, height)
	}
	return table
}

func moveBlizzards(blizzards []grid.Coords, vels []grid.Coords, width int, height int) []grid.Coords {
	newPositions := []grid.Coords{}
	for i, b := range blizzards {
		v := vels[i]
		newPosition := grid.Coords{ X: circularIncrement(b.X, v.X, width), Y: circularIncrement(b.Y, v.Y, height) }
		newPositions = append(newPositions, newPosition)
	}
	return newPositions
}

func circularIncrement(index int, distance int, length int) int {
	return (index + (length + (distance % length))) % length
}

var velocities map[rune]grid.Coords = map[rune]grid.Coords{
	'>': { X: 1, Y: 0},
	'<': { X: -1, Y: 0 },
	'^': { X: 0, Y: -1 },
	'v': { X: 0, Y: 1 },
	' ': { X: 0, Y: 0 },
}
func parseInput(input []string) ([]grid.Coords, []grid.Coords, int, int, grid.Coords, grid.Coords) {
	blizzards, vels := []grid.Coords{}, []grid.Coords{}
	for y, line := range input[1:len(input) - 1] {
		for x, char := range line[1:len(line) - 1] {
			v, isBlizzard := velocities[char]
			if isBlizzard {
				current := grid.Coords{ X: x, Y: y }
				blizzards = append(blizzards, current)
				vels = append(vels, v)
			}
		}
	}
	return blizzards, vels, len(input[0]) - 2, len(input) - 2, grid.Coords{ X: 0, Y: -1 }, grid.Coords{ X: len(input[0]) - 3, Y: len(input) - 2 }
}
