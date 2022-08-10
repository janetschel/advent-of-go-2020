package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/priorityqueue"
	"advent-of-go/utils/sets"
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := files.ReadFile(15, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	grid := parseInput(input)
	return dijkstraPriorityQueue(grid, 1)
}

func solvePart2(input []string) int {
	start := time.Now()
	grid := parseInput(input)
	result := dijkstraPriorityQueue(grid, 5)
	elapsed := time.Since(start)
	fmt.Printf("Part 2 took %s to run\n", elapsed)
	return result
}

func parseInput(input []string) [][]int {
	grid := [][]int{}
	for row := range input {
		grid = append(grid, []int{})
		for col := range input[row] {
			val, _ := strconv.Atoi(input[row][col : col+1])
			grid[row] = append(grid[row], val)
		}
	}
	return grid
}

func dijkstraPriorityQueue(g [][]int, multiplier int) int {
	dist := map[string]int{}

	grid := buildFullGrid(g, multiplier)

	lenX, lenY := len(grid[0]), len(grid) // len(grid[0])*multiplier, len(grid)*multiplier
	target, source := getCoordsKey(lenX-1, lenY-1), getCoordsKey(0, 0)

	dist[source] = 0
	q := make(priorityqueue.PriorityQueue, 0)

	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			v := getCoordsKey(x, y)
			if v != source {
				dist[v] = math.MaxInt
			}
			heap.Push(&q, &priorityqueue.Item{Priority: dist[v], Value: v})
		}
	}

	for q.Len() > 0 {
		u := heap.Pop(&q).(*priorityqueue.Item)
		neighbors := getNeighbors(u.Value, grid, multiplier)
		for _, v := range neighbors {
			if q.Has(v) {
				x, y := toCoords(v)
				alt := dist[u.Value] + grid[y][x]
				if alt < dist[v] {
					dist[v] = alt
					q.Update(v, alt)
				}
			}
		}
	}

	return dist[target]
}

func dijkstra(grid [][]int, multiplier int) int {
	vertices := sets.New()
	dist, prev := map[string]int{}, map[string]string{}

	lenX, lenY := len(grid[0]), len(grid)

	target := getCoordsKey((multiplier*lenX)-1, (multiplier*lenY)-1)

	for row := 0; row < lenY*multiplier; row++ {
		for col := 0; col < lenX*multiplier; col++ {
			key := getCoordsKey(col, row)
			dist[key] = math.MaxInt
			prev[key] = ""
			vertices.Add(key)
		}
	}
	dist["0,0"] = 0

	for vertices.Size() > 0 {
		minDist, minKey := math.MaxInt, ""
		for key, val := range dist {
			if val < minDist && vertices.Has(key) {
				minDist = val
				minKey = key
			}
		}

		vertices.Remove(minKey)
		if minKey == target {
			return dist[target]
		}
		neighbors := getNeighbors(minKey, grid, multiplier)
		for _, key := range neighbors {
			if vertices.Has(key) {
				risk := getRisk(key, grid)
				alt := dist[minKey] + risk
				if alt < dist[key] {
					dist[key] = alt
					prev[key] = minKey
				}
			}
		}
	}

	return dist[target]
}

func getNeighbors(key string, grid [][]int, multiplier int) []string {
	neighbors := []string{}

	x, y := toCoords(key)

	candidates := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, pair := range candidates {
		dx, dy := pair[0], pair[1]
		newX, newY := dx+x, dy+y
		if newX >= 0 && newY >= 0 && !(newX == x && newY == y) && newY < (multiplier*len(grid)) && newX < (multiplier*len(grid[0])) {
			neighbors = append(neighbors, getCoordsKey(newX, newY))
		}
	}

	return neighbors
}

func getCoordsKey(x int, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func toCoords(key string) (int, int) {
	parts := strings.Split(key, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

func getRisk(key string, grid [][]int) int {
	x, y := toCoords(key)
	lenX, lenY := len(grid[0]), len(grid)

	newX, newY := x%lenX, y%lenY
	row, col := y/lenY, x/lenX

	risk := grid[newY][newX]

	for i := 0; i < row+col; i++ {
		risk++
		if risk > 9 {
			risk = 1
		}
	}

	return risk
}

func buildFullGrid(original [][]int, multiplier int) [][]int {
	grid := [][]int{}
	for row := 0; row < len(original)*multiplier; row++ {
		grid = append(grid, []int{})
		for col := 0; col < len(original[0])*multiplier; col++ {
			grid[row] = append(grid[row], getRisk(getCoordsKey(col, row), original))
		}
	}
	return grid
}
