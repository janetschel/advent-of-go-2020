package main

import (
	"sort"
	"strconv"
	"tblue-aoc-2021/utils/files"
)

func main() {
	input := files.ReadFile(9, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	heightMap := parseInput(input)

	return sumLowPoints(heightMap)
}

func solvePart2(input []string) int {
	heightMap := parseInput(input)
	b := buildBasins(heightMap)
	sort.Slice(b, func(i, j int) bool { return len(b[i]) > len(b[j]) })

	return len(b[0]) * len(b[1]) * len(b[2])
}

func parseInput(input []string) [][]int {
	heightMap := [][]int{}
	for _, val := range input {
		row := []int{}
		for _, r := range val {
			v, _ := strconv.Atoi(string(r))
			row = append(row, v)
		}
		heightMap = append(heightMap, row)
	}
	return heightMap
}

func sumLowPoints(heightMap [][]int) int {
	sum := 0
	for i, val := range heightMap {
		for j, h := range val {
			if checkLowPoint(heightMap, i, j) {
				sum += h + 1
			}
		}
	}
	return sum
}

func checkLowPoint(heightMap [][]int, x int, y int) bool {
	h := heightMap[x][y]
	//check above
	if x-1 >= 0 {
		if heightMap[x-1][y] <= h {
			return false
		}
	}
	//check below
	if x+1 < len(heightMap) {
		if heightMap[x+1][y] <= h {
			return false
		}
	}
	//check left
	if y-1 >= 0 {
		if heightMap[x][y-1] <= h {
			return false
		}
	}
	//check right
	if y+1 < len(heightMap[x]) {
		if heightMap[x][y+1] <= h {
			return false
		}
	}
	return true
}

func buildBasins(heightMap [][]int) []map[string]int {
	basins := []map[string]int{}
	for i, val := range heightMap {
		for j := range val {
			if checkLowPoint(heightMap, i, j) {
				pointsHash := map[string]int{}
				pointsHash = expandBasinRecusive(heightMap, pointsHash, i, j)
				basins = append(basins, pointsHash)
			}
		}
	}
	return basins
}

func expandBasinRecusive(heightMap [][]int, basin map[string]int, x int, y int) map[string]int {
	key := strconv.Itoa(x) + "," + strconv.Itoa(y)
	_, ex := basin[key]
	if x < 0 || y < 0 || x >= len(heightMap) || y >= len(heightMap[x]) || heightMap[x][y] == 9 || ex {
		return basin
	}
	basin[key] = 1
	basin = expandBasinRecusive(heightMap, basin, x-1, y)
	basin = expandBasinRecusive(heightMap, basin, x+1, y)
	basin = expandBasinRecusive(heightMap, basin, x, y-1)
	basin = expandBasinRecusive(heightMap, basin, x, y+1)
	return basin

}
