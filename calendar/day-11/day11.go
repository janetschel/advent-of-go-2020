package main

import (
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/slices"
	"advent-of-go-2020/utils/str"
)

type Direction struct {
	deltaRow int
	deltaCol int
}

var known = []Direction {
	{ deltaRow: 0, deltaCol: 1},
	{ deltaRow: 0, deltaCol: -1},
	{ deltaRow: 1, deltaCol: 0},
	{ deltaRow: 1, deltaCol: 1},
	{ deltaRow: 1, deltaCol: -1},
	{ deltaRow: -1, deltaCol: 0},
	{ deltaRow: -1, deltaCol: 1},
	{ deltaRow: -1, deltaCol: -1},
}

func main() {
	input := files.ReadFile(11, "\n")

	solutionPart1 := solvePart1(input)
	solutionPart2 := solvePart2(input)

	println("Occupied seats, part 1 solution:", solutionPart1)
	println("Occupied seats, part 2 solution:", solutionPart2)
}

func solvePart1(input []string) int {
	framed := slices.Frame(input)
	iteration, firstIteration := make([]string, len(framed)), true
	copy(iteration, framed)

	for !slices.Equals(framed, iteration) || firstIteration {
		firstIteration = false
		copy(framed, iteration)

		for row := 1; row < len(framed) - 1; row++ {
			for col := 1; col < len(framed[row]) - 1; col++ {
				next := make([]string, 8)

				for i, element := range known {
					next[i] = string(framed[row + element.deltaRow][col + element.deltaCol])
				}

				count := slices.CountCharInSlice(next, "#")

				if string(framed[row][col]) == "L" && count == 0 {
					iteration[row] = str.ReplaceCharAt(iteration[row], "#", col)
				} else if string(framed[row][col]) == "#" && count >= 4 {
					iteration[row] = str.ReplaceCharAt(iteration[row], "L", col)
				}
			}
		}
	}

	return slices.CountCharInSlice(framed, "#")
}

func solvePart2(input []string) int {
	framed := slices.Frame(input)
	iteration, firstIteration := make([]string, len(framed)), true
	copy(iteration, framed)

	for !slices.Equals(framed, iteration) || firstIteration {
		firstIteration = false
		copy(framed, iteration)

		for row := 1; row < len(framed) - 1; row++ {
			for col := 1; col < len(framed[row]) - 1; col++ {
				next := make([]string, 8)

				for i, element := range known {
					rowMult, colMult := 0, 0
					currentRow := row + element.deltaRow + rowMult
					currentCol := col + element.deltaCol + colMult

					candidate := "."
					for candidate == "." && currentRow >= 0 && currentRow < len(framed) - 1 && currentCol >= 0 && currentCol < len(framed[i]) - 1 {
						candidate = string(framed[currentRow][currentCol])

						rowMult += element.deltaRow
						colMult += element.deltaCol
						currentRow = row + element.deltaRow + rowMult
						currentCol = col + element.deltaCol + colMult
					}

					next[i] = candidate
				}

				count := slices.CountCharInSlice(next, "#")

				if string(framed[row][col]) == "L" && count == 0 {
					iteration[row] = str.ReplaceCharAt(iteration[row], "#", col)
				} else if string(framed[row][col]) == "#" && count >= 5 {
					iteration[row] = str.ReplaceCharAt(iteration[row], "L", col)
				}
			}
		}
	}

	return slices.CountCharInSlice(framed, "#")
}
