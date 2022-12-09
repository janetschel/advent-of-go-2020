package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"strconv"
)

func main() {
	input := files.ReadFile(8, 2022, "\n")
	visible, maxScenicScore := solve(input)
	println(visible)
	println(maxScenicScore)
}

func solve(input []string) (int, int) {
	directions := []grid.Coords{ { X: 0, Y: -1 }, { X: 0, Y: 1 }, { X: -1, Y: 0 }, { X: 1, Y: 0 } }
	trees := parseInput(input)
	visible := grid.PerimeterSize(trees)

	maxScore := 0
	for y := 1; y < len(trees) - 1; y++ {
		for x := 1; x < len(trees[y]) - 1; x++ {
			tree := grid.Coords { X: x, Y: y }
			isVisible := false
			score := 1
			for i := 0; i < len(directions); i++ {
				visibility, viewingDistance := visibilityFromDirection(tree, trees, directions[i])
				isVisible = isVisible || visibility
				score *= viewingDistance
			}
			if isVisible {
				visible++
			}
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return visible, maxScore
}

func parseInput(input []string) [][]int {
	grid := make([][]int, len(input))
	for i, line := range input {
		grid[i] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			height, _ := strconv.Atoi(line[j:j+1])
			grid[i][j] = height
		}
	}
	return grid
}

func visibilityFromDirection(tree grid.Coords, trees [][]int, direction grid.Coords) (bool, int) {
	isVisible, viewingDistance := true, 0

	current := grid.Coords{ X: tree.X + direction.X, Y: tree.Y + direction.Y }
	treeValue := trees[tree.Y][tree.X]
	for isVisible && grid.IsInGrid(current, trees) {
		if trees[current.Y][current.X] >= treeValue {
			isVisible = false
		}
		current.X += direction.X
		current.Y += direction.Y
		viewingDistance++
	}

	return isVisible, viewingDistance
}
