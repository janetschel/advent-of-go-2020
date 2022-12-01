package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"strings"
)

func main() {
	input := files.ReadFile(12, 2017, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	nodes := parseGraph(input)
	visited := walkGraph(nodes, "0", sets.New())
	return visited.Size()
}

func solvePart2(input []string) int {
	nodes := parseGraph(input)
	groups := 0
	startingNode := "0"
	for len(nodes) > 0  && startingNode != "" {
		visited := walkGraph(nodes, startingNode, sets.New())
		groups++
		for _, node := range visited.Iterator() {
			delete(nodes, node)
		}
		for key := range nodes {
			startingNode = key
			break
		}
	}


	return groups
}

func walkGraph(nodes map[string][]string, startingNode string, visited sets.Set) sets.Set {
	if visited.Has(startingNode) {
		return visited
	}
	visited.Add(startingNode)
	for _, child := range nodes[startingNode] {
		walkGraph(nodes, child, visited)
	}
	return visited
}

func parseGraph(input []string) map[string][]string {
	nodes := map[string][]string{}
	for _, line := range input {
		parts := strings.Split(line, " <-> ")
		nodes[parts[0]] = strings.Split(parts[1], ", ")
	}
	return nodes
}
