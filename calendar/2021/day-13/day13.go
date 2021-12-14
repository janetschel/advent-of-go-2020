package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

func main() {
	input := files.ReadFile(13, 2021, "\n\n")
	visible, grid := solvePart1(input)
	println(visible)
	printGrid(grid, "code")
}

func solvePart1(input []string) (int, map[string]bool) {
	grid, folds := parseInput(input)

	visibleAfterFirstFold := 0
	printGrid(grid, "folds/pre")
	for i, f := range folds {
		for key, visible := range grid {
			if visible {
				grid[key] = false
				grid[transformCoords(key, f)] = true
			}
		}
		if i == 0 {
			for _, visible := range grid {
				if visible {
					visibleAfterFirstFold++
				}
			}
		}
		printGrid(grid, fmt.Sprintf("folds/fold-%v", i))
	}

	return visibleAfterFirstFold, grid
}

func parseInput(input []string) (map[string]bool, []coords) {
	grid, folds := map[string]bool{}, []coords{}
	coordsInput := strings.Split(input[0], "\n")
	for i := range coordsInput {
		grid[coordsInput[i]] = true
	}

	foldsInput := strings.Split(input[1], "\n")
	for i := range foldsInput {
		parts := strings.Split(strings.Fields(foldsInput[i])[2], "=")
		if parts[0] == "x" {
			x, _ := strconv.Atoi(parts[1])
			folds = append(folds, coords{x: x, y: 0})
		} else {
			y, _ := strconv.Atoi(parts[1])
			folds = append(folds, coords{x: 0, y: y})
		}
	}
	return grid, folds
}

func transformCoords(coordsKey string, transform coords) string {
	c := toCoords(coordsKey)
	x, y := c.x, c.y
	if transform.x != 0 && x > transform.x {
		x = transform.x - (x - transform.x)
	}
	if transform.y != 0 && y > transform.y {
		y = transform.y - (y - transform.y)
	}
	return fmt.Sprintf("%v,%v", x, y)
}

func getCoordsKey(c coords) string {
	return fmt.Sprintf("%v,%v", c.x, c.y)
}

func toCoords(key string) coords {
	parts := strings.Split(key, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coords{x: x, y: y}
}

func printGrid(points map[string]bool, fileName string) {
	minX, minY, maxX, maxY := dimensions(points)

	lenX, lenY := maxX-minX+1, maxY-minY+1

	img := image.NewRGBA(image.Rect(0, 0, lenX, lenY))
	for key, visible := range points {
		if visible {
			c := toCoords(key)
			img.Set(c.x-minX, c.y-minY, color.White)
		}
	}

	f, err := os.OpenFile(fmt.Sprintf("./calendar/2021/day-13/%v.png", fileName), os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	if err != nil {
		fmt.Errorf("Error opening file", err)
	}
}

func dimensions(points map[string]bool) (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for key, visible := range points {
		if visible {
			c := toCoords(key)
			if c.x > maxX {
				maxX = c.x
			} else if c.x < minX {
				minX = c.x
			}
			if c.y > maxY {
				maxY = c.y
			} else if c.y < minY {
				minY = c.y
			}
		}
	}
	return minX, minY, maxX, maxY
}
