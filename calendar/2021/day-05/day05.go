package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

type pair struct {
	x, y int
}
type line struct {
	start, end pair
}

func main() {
	input := files.ReadFile(05, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	lines := parselines(input)
	// filter out non simple lines
	simplelines := []line{}
	for _, line := range lines {
		if line.start.x == line.end.x || line.start.y == line.end.y {
			simplelines = append(simplelines, line)
		}
	}

	plotMap := createPlotMap(simplelines)

	return countSharedPoints(plotMap)
}

func solvePart2(input []string) int {
	lines := parselines(input)
	plotMap := createPlotMap(lines)
	return countSharedPoints(plotMap)
}

func parselines(input []string) []line {
	lines := []line{}
	for _, val := range input {
		pointsStrings := strings.Split(val, " -> ")
		startPointString := strings.Split(pointsStrings[0], ",")
		endPointString := strings.Split(pointsStrings[1], ",")
		startX, _ := strconv.Atoi(startPointString[0])
		startY, _ := strconv.Atoi(startPointString[1])
		endX, _ := strconv.Atoi(endPointString[0])
		endY, _ := strconv.Atoi(endPointString[1])
		line := line{
			start: pair{
				x: startX,
				y: startY,
			},
			end: pair{
				x: endX,
				y: endY,
			},
		}
		lines = append(lines, line)

	}
	return lines
}

func createPlotMap(lines []line) map[string]int {
	plotMap := make(map[string]int)
	for _, line := range lines {
		if line.start.y == line.end.y {
			//horizontal
			basePoint := line.start
			if line.end.x < line.start.x {
				basePoint = line.end
			}
			for i := 0; i <= int(math.Abs(float64(line.end.x-line.start.x))); i++ {
				currentPoint := pair{
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
				currentPoint := pair{
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
				currentPoint := pair{
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
			count++
		}
	}
	return count
}
