package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"regexp"
	"strconv"
)

type coords struct {
	x int
	y int
}

type targetArea struct {
	min coords
	max coords
}

func main() {
	input := files.ReadFile(17, 2021, "\n")
	target := parseInput(input[0])
	maxY, validVelocities := findBestVelocity(target)
	println(maxY)
	println(validVelocities)
}

func findBestVelocity(target targetArea) (int, int) {
	maxY, validVelocities := 0, 0
	for x := 0; x <= target.max.x; x++ {
		for y := target.min.y; y <= maths.Abs(target.min.y); y++ {
			result, yReached := simulateProbe(coords{x: x, y: y}, target)
			if result {
				validVelocities++
				if yReached > maxY {
					maxY = yReached
				}
			}
		}
	}
	return maxY, validVelocities
}

func parseInput(input string) targetArea {
	numberRegex, _ := regexp.Compile("-?[0-9]+")
	matches := numberRegex.FindAllString(input, -1)
	minX, _ := strconv.Atoi(matches[0])
	maxX, _ := strconv.Atoi(matches[1])
	minY, _ := strconv.Atoi(matches[2])
	maxY, _ := strconv.Atoi(matches[3])
	return targetArea{
		min: coords{x: minX, y: minY},
		max: coords{x: maxX, y: maxY},
	}
}

func simulateProbe(velocity coords, target targetArea) (bool, int) {
	c := coords{x: 0, y: 0}
	maxY := c.y
	for !isInTargetArea(c, target) && !hasMissedTargetArea(c, velocity, target) {
		c, velocity = moveProbe(c, velocity)
		if c.y > maxY {
			maxY = c.y
		}
	}
	return isInTargetArea(c, target), maxY
}

func isInTargetArea(c coords, target targetArea) bool {
	return c.x >= target.min.x && c.x <= target.max.x &&
		c.y >= target.min.y && c.y <= target.max.y
}

func hasMissedTargetArea(c coords, velocity coords, target targetArea) bool {
	if velocity.y > 0 {
		return false
	}
	return c.x > target.max.x || c.y < target.min.y
}

func moveProbe(c coords, velocity coords) (coords, coords) {
	newCoords := coords{x: c.x + velocity.x, y: c.y + velocity.y}
	newVelocity := coords{x: velocity.x - 1, y: velocity.y - 1}
	if newVelocity.x < 0 {
		newVelocity.x = 0
	}
	return newCoords, newVelocity
}
