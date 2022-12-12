package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type point struct {
	position *grid.Coords
	velocity *grid.Coords
}

func main() {
	input := files.ReadFile(10, 2018, "\n")
	solve(input)
}

func solve(input []string) {
	points := parseInput(input)
	// this killswitch number was picked after determining when the message appears
	// and is not guaranteed to work with all inputs
	for killswitch := 1; killswitch < 12000; killswitch++ {
		moveAllPoints(points)
		gridString := buildReasonableString(points)
		if gridString != "" {
			println(killswitch, "seconds")
			println(gridString)
		}
	}

	println("see logs for answers to part one and two")
}

func parseInput(input []string) []*point {
	points := make([]*point, len(input))
	number := regexp.MustCompile("-?[0-9]+")

	for i, line := range input {
		numbers := number.FindAllString(line, -1)
		px, _ := strconv.Atoi(numbers[0])
		py, _ := strconv.Atoi(numbers[1])
		vx, _ := strconv.Atoi(numbers[2])
		vy, _ := strconv.Atoi(numbers[3])
		pos, vel := grid.Coords{ X: px, Y: py}, grid.Coords{ X: vx, Y: vy }
		p := point{ position: &pos, velocity: &vel }
		points[i] = &p
	}

	return points
}

func buildReasonableString(points []*point) string {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	pointSet := sets.New()

	for _, p := range points {
		if p.position.X < minX {
			minX = p.position.X
		}

		if p.position.X > maxX {
			maxX = p.position.X
		}

		if p.position.Y < minY {
			minY = p.position.Y
		}

		if p.position.Y > maxY {
			maxY = p.position.Y
		}

		pointSet.Add(p.position.ToString())
	}

	gridString := ""
	// this size bound was selected after revealing the message
	// and is not guaranteed to work for all inputs
	if maths.Abs(maxX - minX) <= 100 && maths.Abs(maxY - minY) <= 100 {
		for y := minY; y <= maxY; y++ {
			line := ""
			for x := minX; x <= maxX; x++ {
				key := fmt.Sprintf("%d,%d", x, y)
				if pointSet.Has(key) {
					line += "█"
				} else {
					line += " "
				}
			}
			gridString += line + "\n"
		}
	}
	return gridString
}

func moveAllPoints(points []*point) {
	for _, p := range points {
		p.move()
	}
}

func (p *point) move() {
	p.position.X += p.velocity.X
	p.position.Y += p.velocity.Y
}
