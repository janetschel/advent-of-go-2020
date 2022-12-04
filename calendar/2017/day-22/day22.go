package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"fmt"
)

type nodeState int
const (
	clean nodeState = 1
	weakened = 2
	infected = 3
	flagged = 4
)

func main() {
	input := files.ReadFile(22, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	nodeStates, origin := parseInput(input)
	nextStates := map[nodeState]func(nodeState, grid.Coords) (nodeState, grid.Coords){
		clean: doInfect,
		infected: doClean,
	}
	return runBursts(nodeStates, origin, nextStates, 10000)
}

func solvePart2(input []string) int {
	nodeStates, origin := parseInput(input)
	nextStates := map[nodeState]func(nodeState, grid.Coords) (nodeState, grid.Coords){
		clean: updateClean,
		weakened: updateWeakened,
		infected: updateInfected,
		flagged: updateFlagged,
	}
	return runBursts(nodeStates, origin, nextStates, 10000000)
}

func runBursts(nodeStates map[string]nodeState, origin grid.Coords, nextState map[nodeState]func(nodeState, grid.Coords) (nodeState, grid.Coords), burstCount int) int {
	current, velocity := grid.Coords{ X: origin.X, Y: origin.Y}, grid.Coords{ X: 0, Y: 1}
	infectionBursts := 0

	for i := 0; i < burstCount; i++ {
		key := coordsKey(current)
		state, hasState := nodeStates[key]
		if !hasState {
			state = clean
		}
		var newState nodeState
		newState, velocity = nextState[state](state, velocity)
		nodeStates[key] = newState
		if newState == infected {
			infectionBursts++
		}
		current.X += velocity.X
		current.Y += velocity.Y
	}

	return infectionBursts
}

func coordsKey(coords grid.Coords) string {
	return fmt.Sprintf("%d,%d", coords.X, coords.Y)
}

func turnRight(velocity grid.Coords) grid.Coords {
	return grid.RotateCoordsCounterclockwise(velocity, grid.Origin, float64(-90))
}

func turnLeft(velocity grid.Coords) grid.Coords {
	return grid.RotateCoordsCounterclockwise(velocity, grid.Origin, float64(90))
}

func parseInput(input []string) (map[string]nodeState, grid.Coords) {
	originX, originY := len(input[0]) / 2, len(input) / 2
	nodeStates := map[string]nodeState{}

	for y, row := range input {
		for x, value := range row {
			coords := grid.Coords{ X: x, Y: -1 * y }
			key := coordsKey(coords)
			if value == '#' {
				nodeStates[key] = infected
			} else {
				nodeStates[key] = clean
			}
		}
	}

	return nodeStates, grid.Coords{ X: originX, Y: -1 * originY}
}

func updateClean(state nodeState, velocity grid.Coords) (nodeState, grid.Coords) {
	return weakened, turnLeft(velocity)
}

func updateWeakened(state nodeState, velocity grid.Coords) (nodeState, grid.Coords) {
	return infected, grid.Coords{ X: velocity.X, Y: velocity.Y }
}

func updateInfected(state nodeState, velocity grid.Coords) (nodeState, grid.Coords) {
	return flagged, turnRight(velocity)
}

func updateFlagged(state nodeState, velocity grid.Coords) (nodeState, grid.Coords) {
	return clean, grid.Coords{ X: -1 * velocity.X, Y: -1 * velocity.Y }
}

func doInfect(state nodeState, velocity grid.Coords) (nodeState, grid.Coords) {
	return infected, turnLeft(velocity)
}

func doClean(state nodeState, velocity grid.Coords) (nodeState, grid.Coords) {
	return clean, turnRight(velocity)
}
