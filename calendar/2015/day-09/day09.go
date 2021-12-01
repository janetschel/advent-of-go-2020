package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(9, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	matrix := buildMatrix(input)
	return findShortestPath(matrix)
}

func solvePart2(input []string) int {
	matrix := buildMatrix(input)
	return findLongestPath(matrix)
}

func findShortestPath(matrix [][]int) int {
	indexes := make([]int, len(matrix))
	for i := range indexes {
		indexes[i] = i
	}
	permutations := buildPermutations(indexes)
	
	min := math.MaxInt
	for _, path := range permutations {
		length := calculatePathLength(path, matrix)
		if (length < min) {
			min = length
		}
	}
	return min
}

func findLongestPath(matrix [][]int) int {
	indexes := make([]int, len(matrix))
	for i := range indexes {
		indexes[i] = i
	}
	permutations := buildPermutations(indexes)
	
	max := 0
	for _, path := range permutations {
		length := calculatePathLength(path, matrix)
		if (length > max) {
			max = length
		}
	}
	return max
}

func buildMatrix(input []string) [][]int {
	nodes := buildNodeList(input)
	matrix := make([][]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		matrix[i] = make([]int, len(nodes))
		for j := 0; j < len(nodes); j++ {
				if (i == j) {
					matrix[i][j] = 0
				} else {
				matchingEdges := slices.Filter(input, func(line string) bool {
					return strings.Contains(line, nodes[i] + " to " + nodes[j]) || strings.Contains(line, nodes[j] + " to " + nodes[i])
				})
				if len(matchingEdges) > 0 {
					distance, _ := strconv.Atoi(strings.Split(matchingEdges[0], " = ")[1])
					matrix[i][j] = distance
				} else {
					matrix[i][j] = math.MaxInt
				}
			}
		}
	}
	return matrix
}

func buildNodeList(input []string) []string {
	nodeMap := make(map[string]bool)
	for i := range input {
		lineParts := strings.Split(input[i], " = ")
		cities := strings.Split(lineParts[0], " to ")
		nodeMap[cities[0]] = true
		nodeMap[cities[1]] = true
	}
	nodes := []string {}
	for key := range nodeMap {
		nodes = append(nodes, key)
	}
	return nodes
}

func buildPermutations(nodes []int) [][]int {
	length := len(nodes)
	initial := make([]int, length)
	copy(initial, nodes)
	generatedPermutations := [][]int { initial }
	indexes := make([]int, length)

	i := 0
	for i < length {
		if indexes[i] < i {
			if i % 2 == 0 {
				slices.Swap(nodes, 0, i)
			} else {
				slices.Swap(nodes, indexes[i], i)
			}
			perm := make([]int, length)
			copy(perm, nodes)
			generatedPermutations = append(generatedPermutations, perm)
			indexes[i] = indexes[i] + 1
			i = 0
		} else {
			indexes[i] = 0
			i++
		}
	}

	return generatedPermutations
}

func calculatePathLength(path []int, matrix [][]int) int {
	length := 0
	for i := 0; i < len(path) - 1; i++ {
		length += matrix[path[i]][path[i + 1]]
	}
	if (length < 200) {
		fmt.Printf("%v %v\n", length, path)
	}
	return length
}
