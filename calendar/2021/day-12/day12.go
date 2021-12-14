package main

import (
	"advent-of-go/utils/files"
	"strings"
)

type edge struct {
	start string
	end   string
}

func main() {
	input := files.ReadFile(12, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	graph := parseInput(input)
	paths := [][]string{}

	buildAllPaths("start", "end", graph, &paths, 1)
	return len(paths)
}

func solvePart2(input []string) int {
	graph := parseInput(input)
	paths := [][]string{}

	buildAllPaths("start", "end", graph, &paths, 2)
	return len(paths)
}

func buildAllPaths(start string, end string, graph []edge, paths *[][]string, maxSmallVisits int) {
	visited := map[string]int{}
	path := []string{}

	buildAllPathsFromCurrent(start, end, visited, path, graph, paths, maxSmallVisits)
}

func buildAllPathsFromCurrent(current string, end string, visited map[string]int, path []string, graph []edge, paths *[][]string, maxSmallVisits int) {
	if isSmall(current) {
		visited[current]++
	}

	maxVisits := maxSmallVisits
	for cave, visits := range visited {
		if isSmall(cave) && visits >= maxSmallVisits {
			maxVisits = 1
		}
	}
	path = append(path, current)

	if current == end {
		*paths = append(*paths, path)
	} else {
		adjacent := getAdjacent(current, graph)
		for _, adj := range adjacent {
			if visited[adj] < maxVisits && adj != "start" {
				buildAllPathsFromCurrent(adj, end, visited, path, graph, paths, maxVisits)
			}
		}
	}

	path = path[:len(path)-1]
	if isSmall(current) {
		visited[current]--
	}
}

func getAdjacent(current string, graph []edge) []string {
	adj := []string{}
	for i := range graph {
		if graph[i].start == current {
			adj = append(adj, graph[i].end)
		}
		if graph[i].end == current {
			adj = append(adj, graph[i].start)
		}
	}
	return adj
}

func isSmall(cave string) bool {
	return cave[:1] == strings.ToLower(cave[:1])
}

func parseInput(input []string) []edge {
	edges := []edge{}
	for i := range input {
		parts := strings.Split(input[i], "-")
		e := edge{
			start: parts[0],
			end:   parts[1],
		}
		edges = append(edges, e)
	}
	return edges
}
