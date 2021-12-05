package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

type Pair struct {
	x, y int
}
type Line struct {
	start, end Pair
}

func main() {
	input := files.ReadFile(05, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	lines := parseLines(input)
	// filter out non simple lines
	simpleLines := []Line{}
	for _, line := range lines {
		if line.start.x == line.end.x || line.start.y == line.end.y {
			simpleLines = append(simpleLines, line)
		}
	}

	plotMap := createPlotMap(simpleLines)

	return countSharedPoints(plotMap)
}

func solvePart2(input []string) int {
	lines := parseLines(input)
	plotMap := createPlotMap(lines)
	return countSharedPoints(plotMap)
}

func parseLines(input []string) []Line {
	lines := []Line{}
	for _, val := range input {
		pointsStrings := strings.Split(val, " -> ")
		startPointString := strings.Split(pointsStrings[0], ",")
		endPointString := strings.Split(pointsStrings[1], ",")
		startX, _ := strconv.Atoi(startPointString[0])
		startY, _ := strconv.Atoi(startPointString[1])
		endX, _ := strconv.Atoi(endPointString[0])
		endY, _ := strconv.Atoi(endPointString[1])
		line := Line{
			start: Pair{
				x: startX,
				y: startY,
			},
			end: Pair{
				x: endX,
				y: endY,
			},
		}
		lines = append(lines, line)

	}
	return lines
}

func createPlotMap(lines []Line) map[string]int {
	plotMap := make(map[string]int)
	for _, line := range lines {
		if line.start.y == line.end.y {
			//horizontal
			basePoint := line.start
			if line.end.x < line.start.x {
				basePoint = line.end
			}
			for i := 0; i <= int(math.Abs(float64(line.end.x-line.start.x))); i++ {
				currentPoint := Pair{
					x: basePoint.x + i,
					y: basePoint.y,
				}
				pointString := fmt.Sprintf("%v", currentPoint)
				val := plotMap[pointString]
				plotMap[pointString] = val + 1
			}
		} else if line.start.x == line.end.x {
			basePoint := line.start
			//vertical
			if line.end.y < line.start.y {
				basePoint = line.end
			}
			for i := 0; i <= int(math.Abs(float64(line.end.y-line.start.y))); i++ {
				currentPoint := Pair{
					x: basePoint.x,
					y: basePoint.y + i,
				}
				pointString := fmt.Sprintf("%v", currentPoint)
				val := plotMap[pointString]
				plotMap[pointString] = val + 1
			}
		} else {
			xVal, yVal := 1, 1
			if line.end.x < line.start.x {
				xVal = -1
			}
			if line.end.y < line.start.y {
				yVal = -1
			}
			for i := 0; i <= int(math.Abs(float64(line.end.x-line.start.x))); i++ {
				currentPoint := Pair{
					x: line.start.x + (i * xVal),
					y: line.start.y + (i * yVal),
				}
				pointString := fmt.Sprintf("%v", currentPoint)
				val := plotMap[pointString]
				plotMap[pointString] = val + 1
			}
		}
		//diagonal
	}
	return plotMap
}

func countSharedPoints(m map[string]int) int {
	count := 0
	for _, value := range m {
		if value > 1 {
			count += 1
		}
	}
	return count
}
