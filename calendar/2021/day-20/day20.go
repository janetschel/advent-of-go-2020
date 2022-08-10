package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coords struct{ x, y int }
type image map[string]string

func main() {
	input := files.ReadFile(20, 2021, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	algo, imgSet := parseMap(input)
	enhancedSet := enhanceSteps(imgSet, algo, 2, true)
	return countLights(enhancedSet)
}

func solvePart2(input []string) int {
	algo, imgSet := parseMap(input)
	enhancedSet := enhanceSteps(imgSet, algo, 50, false)
	return countLights(enhancedSet)
}

func parseMap(input []string) (string, image) {
	lines := strings.Split(input[1], "\n")
	img := map[string]string{}
	for y, line := range lines {
		for x, char := range line {
			key := getCoordsKey(coords{x: x, y: y})
			img[key] = string(char)
		}
	}
	return input[0], img
}

func enhanceSteps(img image, algo string, steps int, withDebugging bool) image {
	enhanced := img
	borders := []string{".", "#"}
	if algo[0:1] == "." {
		borders[1] = "."
	}
	if withDebugging {
		printImage(img)
		fmt.Println()
	}
	for i := 1; i <= steps; i++ {
		enhanced = enhance(enhanced, algo, borders[(i-1)%2])
		if withDebugging {
			fmt.Printf("after %v steps\n", i)
			printImage(enhanced)
			fmt.Println()
		}
	}
	return enhanced
}

func getValue(c coords, img image, borderChar string) int {
	bin := ""
	for y := c.y - 1; y <= c.y+1; y++ {
		for x := c.x - 1; x <= c.x+1; x++ {
			key := getCoordsKey(coords{x: x, y: y})
			char, inMap := img[key]
			if !inMap {
				char = borderChar
			}
			if char == "#" {
				bin += "1"
			} else {
				bin += "0"
			}
		}
	}
	value, _ := strconv.ParseInt(bin, 2, 64)
	return int(value)
}

func enhance(img image, algo string, borderChar string) image {
	enhanced := image{}
	min, max := getBounds(img)

	for y := min.y - 3; y <= max.y+3; y++ {
		for x := min.x - 3; x <= max.y+3; x++ {
			c := coords{x: x, y: y}
			value := getValue(c, img, borderChar)
			char := algo[value : value+1]
			enhanced[getCoordsKey(c)] = char
		}
	}
	return enhanced
}

func toCoords(key string) coords {
	parts := strings.Split(key, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coords{x: x, y: y}
}

func getCoordsKey(c coords) string {
	return fmt.Sprintf("%v,%v", c.x, c.y)
}

func getBounds(img image) (coords, coords) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, 0, 0
	for key := range img {
		c := toCoords(key)
		if c.x < minX {
			minX = c.x
		} else if c.x > maxX {
			maxX = c.x
		}
		if c.y < minY {
			minY = c.y
		} else if c.y > maxY {
			maxY = c.y
		}
	}
	return coords{x: minX, y: minY}, coords{x: maxX, y: maxY}
}

func printImage(img image) {
	image := [][]string{}
	min, max := getBounds(img)
	for y := min.y; y <= max.y; y++ {
		image = append(image, make([]string, max.x-min.x+1))
		for x := min.x; x <= max.x; x++ {
			key := getCoordsKey(coords{x: x, y: y})
			char := img[key]
			image[y-min.y][x-min.x] = char
		}
	}
	for i := range image {
		fmt.Println(strings.Join(image[i], ""))
	}
}

func countLights(img image) int {
	lights := 0
	for _, char := range img {
		if char == "#" {
			lights++
		}
	}
	return lights
}
