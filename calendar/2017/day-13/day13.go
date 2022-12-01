package main

import (
	"advent-of-go/utils/files"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(13, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	severity, _ := simulateTrip(input, 0, false)
	return severity
}

func simulateTrip(input []string, delay int, shortCircuit bool) (int, int) {
	firewall := parseInput(input)
	severity, timesCaught := 0, 0

	for layer, layerRange := range firewall {
		position := calculateCurrentPosition(layerRange, layer + delay)
		if position == 0 {
			severity += (layer * layerRange)
			timesCaught++
			if shortCircuit {
				return severity, timesCaught
			}
		}
	}
	return severity, timesCaught
}

func solvePart2(input []string) int {
	timesCaught := math.MaxInt

	for delay := 0; timesCaught != 0; delay++ {
		_, timesCaught = simulateTrip(input, delay, true)
		if timesCaught == 0 {
			return delay
		}
	}

	return math.MaxInt
}

func parseInput(input []string) (map[int]int) {
	firewall := map[int]int{}
	for _, line := range input {
		parts := strings.Split(line, ": ")
		layer, _ := strconv.Atoi(parts[0])
		layerRange, _ := strconv.Atoi(parts[1])
		firewall[layer] = layerRange
	}
	return firewall
}

func calculateCurrentPosition(layerRange int, time int) int {
	offset := time % ((layerRange - 1) * 2)
	if offset > layerRange - 1 {
		return 2 * (layerRange - 1) - offset
	}
	return offset
}
