package main

import (
	"advent-of-go/utils/files"
	"math"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type lineSegment struct {
	start coordinates
	end   coordinates
}

func main() {
	input := files.ReadFile(05, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	lines := parseInput(input)
	horizontalAndVerticalLines := []lineSegment{}
	for i := range lines {
		current := lines[i]
		if current.IsHorizontalOrVertical() {
			horizontalAndVerticalLines = append(horizontalAndVerticalLines, current)
		}
	}
	grid := plotLines(horizontalAndVerticalLines)
	overlaps := 0
	for _, count := range grid {
		if count > 1 {
			overlaps++
		}
	}

	return overlaps
}

func solvePart2(input []string) int {
	lines := parseInput(input)
	grid := plotLines(lines)
	overlaps := 0
	for _, count := range grid {
		if count > 1 {
			overlaps++
		}
	}
	return overlaps
}

func plotLines(lines []lineSegment) map[string]int {
	gridMap := map[string]int{}
	for i := range lines {
		current := lines[i]
		if current.IsVertical() {
			plotVerticalLine(current, gridMap)
		} else {
			minX, maxX := current.start.x, current.end.x
			y := float64(current.start.y)
			if current.end.x < current.start.x {
				minX, maxX = current.end.x, current.start.x
				y = float64(current.end.y)
			}
			slope := current.GetSlope()
			for x := minX; x <= maxX; x++ {
				key := strconv.Itoa(x) + "," + strconv.Itoa(int(y))
				gridMap[key]++
				y += slope
			}
		}
	}
	return gridMap
}

func plotVerticalLine(line lineSegment, gridMap map[string]int) {
	yInc := 1
	if line.start.y >= line.end.y {
		yInc = -1
	}
	for y := line.start.y; y != line.end.y+yInc; y += yInc {
		key := strconv.Itoa(line.start.x) + "," + strconv.Itoa(y)
		gridMap[key]++
	}
}

func parseInput(input []string) []lineSegment {
	lines := []lineSegment{}
	for i := range input {
		lines = append(lines, parseSegment(input[i]))
	}
	return lines
}

func parseSegment(input string) lineSegment {
	parts := strings.Split(input, " -> ")
	return lineSegment{
		start: parseCoords(parts[0]),
		end:   parseCoords(parts[1]),
	}
}

func parseCoords(input string) coordinates {
	parts := strings.Split(input, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coordinates{x: x, y: y}
}

func (line *lineSegment) IsHorizontalOrVertical() bool {
	return line.IsVertical() || line.start.y == line.end.y
}

func (line *lineSegment) IsVertical() bool {
	return line.start.x == line.end.x
}

func (line *lineSegment) GetSlope() float64 {
	if line.IsVertical() {
		return math.NaN()
	}
	return (float64(line.end.y) - float64(line.start.y)) / (float64(line.end.x) - float64(line.start.x))
}
