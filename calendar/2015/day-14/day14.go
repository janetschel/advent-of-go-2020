package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

type Reindeer struct {
	name string
	speed int
	flightDuration int
	restDuration int
	currentTime int
	currentDistance int
	distances []int
}

type StatefulReindeer interface {
	move(seconds ...int)
	traveledDistance(seconds int) int
}

func main() {
	input := files.ReadFile(14, 2015, "\n")
	reindeer := parseInput(input)
	raceReindeer(reindeer, 2503)
	println(solvePart1(reindeer))
	println(solvePart2(reindeer))
}

func solvePart1(reindeer []Reindeer) int {	
	max := 0
	for i := range reindeer {
		if reindeer[i].currentDistance > max {
			max = reindeer[i].currentDistance
		}
	}

	return max
}

func solvePart2(reindeer []Reindeer) int {
	points := make([]int, len(reindeer))

	duration := len(reindeer[0].distances)
	for seconds := 0; seconds < duration; seconds++ {
		maxDistance := 0
		for i := range reindeer {
			currentDistance := reindeer[i].distances[seconds]
			if currentDistance > maxDistance {
				maxDistance = currentDistance
			}
		}
		for i := range reindeer {
			if reindeer[i].distances[seconds] == maxDistance {
				points[i]++
			}
		}
	}

	return slices.Max(points)
}

func raceReindeer(reindeer []Reindeer, raceDuration int) {
	for i := range reindeer {
		reindeer[i].move(raceDuration)
	}
}

func parseReindeer(line string) Reindeer {
	parts := strings.Split(line, " ")
	speed, _ := strconv.Atoi(parts[3])
	flightDuration, _ := strconv.Atoi(parts[6])
	restDuration, _ := strconv.Atoi(parts[13])

	return Reindeer{
		name: parts[0],
		speed: speed,
		flightDuration: flightDuration,
		restDuration: restDuration,
	}
}

func parseInput(input []string) []Reindeer {
	reindeer := make([]Reindeer, len(input))
	for i := range input {
		reindeer[i] = parseReindeer(input[i])
	}
	return reindeer
}

func (reindeer *Reindeer) move(seconds int) {
	for i := 0; i < seconds; i++ {
		reindeer.currentTime++

		if (reindeer.currentTime - 1) % (reindeer.restDuration + reindeer.flightDuration) < reindeer.flightDuration {
			reindeer.currentDistance += reindeer.speed
		}

		reindeer.distances = append(reindeer.distances, reindeer.currentDistance)
	}
}
