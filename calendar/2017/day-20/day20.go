package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type threeDCoords struct {
	x int
	y int
	z int
}
type particle struct {
	position *threeDCoords
	velocity *threeDCoords
	acceleration *threeDCoords
}

func main() {
	input := files.ReadFile(20, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	particles := parseInput(input)

	iterations := 1000
	minDist, minIndex := math.MaxInt, 0
	for t := 0; t < iterations; t++ {
		for i, p := range particles {
			tick(&p)
			if t == iterations - 1 {
				dist := distanceFromOrigin(p)
				if dist < minDist {
					minDist, minIndex = dist, i
				}
			}
		}
	}

	return minIndex
}

func solvePart2(input []string) int {
	particles := parseInput(input)
	particlesSet := sets.New()

	for i := range particles {
		particlesSet.Add(fmt.Sprint(i))
	}


	iterations := 1000
	for t := 0; t < iterations; t++ {
		posistionSet := sets.New()
		originalIndex := map[string]int{}
		for i, p := range particles {
			if particlesSet.Has(fmt.Sprint(i)) {
				tick(&p)
				pos := fmt.Sprintf("%d,%d,%d", p.position.x, p.position.y, p.position.z)
				if posistionSet.Has(pos) {
					particlesSet.Remove(fmt.Sprint(i))
					particlesSet.Remove(fmt.Sprint(originalIndex[pos]))
				} else {
					posistionSet.Add(pos)
					originalIndex[pos] = i
				}
			}
		}
	}

	return particlesSet.Size()
}

func parseInput(input []string) []particle {
	particles := []particle{}

	for _, def := range input {
		parts := strings.Split(def, ", ")
		p := particle{
			position: parseCoord(parts[0]),
			velocity: parseCoord(parts[1]),
			acceleration: parseCoord(parts[2]),
		}
		particles = append(particles, p)
	}

	return particles
}

func parseCoord(input string) *threeDCoords {
	parts := strings.Split(input[3:len(input) - 1], ",")
	x, _ := strconv.Atoi(strings.Trim(parts[0], " "))
	y, _ := strconv.Atoi(strings.Trim(parts[1], " "))
	z, _ := strconv.Atoi(strings.Trim(parts[2], " "))
	coords := threeDCoords{ x: x, y: y, z: z }
	return &coords
}

func distanceFromOrigin(p particle) int {
	return int(
		math.Abs(float64(p.position.x)) +
		math.Abs(float64(p.position.y)) +
		math.Abs(float64(p.position.z)))
}

func tick(p *particle) {
	p.velocity.x += p.acceleration.x
	p.velocity.y += p.acceleration.y
	p.velocity.z += p.acceleration.z

	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	p.position.z += p.velocity.z
}

func print(p particle) {
	println(fmt.Sprintf("p= %d,%d,%d v= %d,%d,%d a= %d,%d,%d", p.position.x, p.position.y, p.position.z, p.velocity.x, p.velocity.y, p.velocity.z, p.acceleration.x, p.acceleration.y, p.acceleration.z))
}
