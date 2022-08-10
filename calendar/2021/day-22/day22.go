package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"fmt"
	"strconv"
	"strings"
)

type step struct {
	turnOff bool
	x       coordRange
	y       coordRange
	z       coordRange
}

type coordRange struct {
	min int
	max int
}

func main() {
	input := files.ReadFile(22, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	steps := parseInput(input)
	return runRebootStepsInInitializationRegion(steps)
}

func solvePart2(input []string) int {
	steps := parseInput(input)
	return runStepsRecursively(steps)
}

func runRebootStepsInInitializationRegion(steps []step) int {
	core := sets.New()
	for _, s := range steps {
		if isInInitializationRegion(s) {
			for x := s.x.min; x <= s.x.max; x++ {
				for y := s.y.min; y <= s.y.max; y++ {
					for z := s.z.min; z <= s.z.max; z++ {
						if !s.turnOff {
							core.Add(getCoordsKey(x, y, z))
						} else {
							core.Remove(getCoordsKey(x, y, z))
						}
					}
				}
			}
		}
	}
	return core.Size()
}

func getCoordsKey(x int, y int, z int) string {
	return fmt.Sprintf("%v,%v,%v", x, y, z)
}

func isInInitializationRegion(s step) bool {
	return s.x.min >= -50 && s.x.max <= 50 &&
		s.y.min >= -50 && s.y.max <= 50 &&
		s.z.min >= -50 && s.z.min <= 50
}

func parseInput(input []string) []step {
	steps := []step{}
	for i := range input {
		parts := strings.Split(input[i], ",")
		s := step{}
		if parts[0][:3] == "off" {
			s.turnOff = true
		} else {
			s.turnOff = false
		}
		xParts := strings.Split(strings.Split(parts[0], "=")[1], "..")
		yParts := strings.Split(strings.Split(parts[1], "=")[1], "..")
		zParts := strings.Split(strings.Split(parts[2], "=")[1], "..")

		xMin, _ := strconv.Atoi(xParts[0])
		xMax, _ := strconv.Atoi(xParts[1])
		yMin, _ := strconv.Atoi(yParts[0])
		yMax, _ := strconv.Atoi(yParts[1])
		zMin, _ := strconv.Atoi(zParts[0])
		zMax, _ := strconv.Atoi(zParts[1])

		s.x = coordRange{min: xMin, max: xMax}
		s.y = coordRange{min: yMin, max: yMax}
		s.z = coordRange{min: zMin, max: zMax}

		steps = append(steps, s)
	}
	return steps
}

func getVolume(s step) int {
	return (s.x.max - s.x.min + 1) * (s.y.max - s.y.min + 1) * (s.z.max - s.z.min + 1)
}

func (s1 step) intersection(s2 step) (bool, step) {
	x, X, y, Y, z, Z := s1.x.min, s1.x.max, s1.y.min, s1.y.max, s1.z.min, s1.z.max

	if s2.x.min >= x {
		x = s2.x.min
	}
	if s2.y.min >= y {
		y = s2.y.min
	}
	if s2.z.min >= z {
		z = s2.z.min
	}

	if s2.x.max <= X {
		X = s2.x.max
	}
	if s2.y.max <= Y {
		Y = s2.y.max
	}
	if s2.z.max <= Z {
		Z = s2.z.max
	}

	if x <= X && y <= Y && z <= Z {
		return true, step{x: coordRange{min: x, max: X}, y: coordRange{min: y, max: Y}, z: coordRange{min: z, max: Z}}
	}
	return false, step{}
}

func runStepsRecursively(steps []step) int {
	if len(steps) == 0 {
		return 0
	}

	head, tail := steps[0], steps[1:]
	if head.turnOff {
		return runStepsRecursively(tail)
	}

	intersections := []step{}
	for _, t := range tail {
		doesIntersect, i := head.intersection(t)
		if doesIntersect {
			intersections = append(intersections, i)
		}
	}

	return (getVolume(head) + runStepsRecursively(tail)) - runStepsRecursively(intersections)
}
