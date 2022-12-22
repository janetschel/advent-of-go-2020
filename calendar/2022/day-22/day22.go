package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"advent-of-go/utils/maths"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type puzzle struct {
	board [][]string
	sideLength int
	start grid.Coords
	instructions []string
}

func main() {
	input := files.ReadFile(22, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

var columnMultiplier, rowMultiplier = 4, 1000
func solvePart1(input []string) int {
	finalPosition, direction := traverse(parseInput(input), false)
	return (rowMultiplier * (finalPosition.Y + 1)) + (columnMultiplier * (finalPosition.X + 1)) + direction
}

func solvePart2(input []string) int {
	finalPosition, direction := traverse(parseInput(input), true)
	return (rowMultiplier * (finalPosition.Y + 1)) + (columnMultiplier * (finalPosition.X + 1)) + direction
}

var directions []grid.Coords = []grid.Coords{ { X: 1, Y: 0 }, { X: 0, Y: 1 }, { X: -1, Y: 0 }, { X: 0, Y: -1 }}
var right, left, empty, wall, free = "R", "L", " ", "#", "."
func traverse(p puzzle, treatAsCube bool) (grid.Coords, int) {
	current := grid.Coords{ X: p.start.X, Y: p.start.Y }
	direction := 0
	board, instructions := p.board, p.instructions

	for _, instruction := range instructions {
		if instruction == right {
			direction = turnRight(direction)
		} else if instruction == left {
			direction = turnLeft(direction)
		} else {
			distance, _ := strconv.Atoi(instruction)
			move:
				for steps := 0; steps < distance; {
					next, nextDir := grid.Coords{ X: current.X, Y: current.Y }, direction
					if directions[direction].X < 0 {
						next.X = circularBackwards(current.X, directions[direction].X, len(board[current.Y]))
					} else {
						next.X = circularForwards(current.X, directions[direction].X, len(board[current.Y]))
					}
					if directions[direction].Y < 0 {
						next.Y = circularBackwards(current.Y, directions[direction].Y, len(board))
					} else {
						next.Y = circularForwards(current.Y, directions[direction].Y, len(board))
					}
					for board[next.Y][next.X] == empty {
						if treatAsCube {
							next, nextDir = getNextPositionAndDirectionOnCube(current, direction, p)
						} else {
							if directions[direction].X < 0 {
								next.X = circularBackwards(next.X, directions[direction].X, len(board[next.Y]))
							} else {
								next.X = circularForwards(next.X, directions[direction].X, len(board[next.Y]))
							}
							if directions[direction].Y < 0 {
								next.Y = circularBackwards(next.Y, directions[direction].Y, len(board))
							} else {
								next.Y = circularForwards(next.Y, directions[direction].Y, len(board))
							}
						}
					}
					if board[next.Y][next.X] == wall {
						break move
					}
					steps++
					current.X = next.X
					current.Y = next.Y
					direction = nextDir
				}
		}
	}

	return current, direction
}

func circularBackwards(index int, distance int, length int) int {
	return (index + (length + (distance % length))) % length
}

func circularForwards(index int, distance int, length int) int {
	return (index + distance) % length
}

func turnRight(direction int) int {
	return circularForwards(direction, 1, len(directions))
}

func turnLeft(direction int) int {
	return circularBackwards(direction, -1, len(directions))
}

func parseInput(input []string) puzzle {
	board := make([][]string, len(input) - 2)
	var instructions []string

	maxLength, sideLength := 0, math.MaxInt
	for _, line := range input[:len(input) - 2] {
		maxLength = maths.Max(maxLength, len(line))
		sideLength = maths.Min(sideLength, strings.Count(line, free) + strings.Count(line, wall))
	}

	for i, line := range input {
		if i < len(input) - 2 {
			board[i] = strings.Split(fmt.Sprintf("%-*s", maxLength, line), "")
		} else if i == len(input) - 1 {
			withSpace := strings.ReplaceAll(line, right, " R ")
			withSpace = strings.ReplaceAll(withSpace, left, " L ")
			instructions = strings.Fields(withSpace)
		}
	}

	return puzzle{
		board: board,
		start: grid.Coords{ X: strings.Index(input[0], free), Y: 0 },
		instructions: instructions,
		sideLength: sideLength,
	}
}


// only works for cube nets matching this format
//  12
//  3
// 45
// 6
func getSide(current grid.Coords, p puzzle) int {
	row := current.Y / p.sideLength
	column := current.X / p.sideLength
	if row == 0 && column == 1 {
		return 1
	}
	if row == 0 && column == 2 {
		return 2
	}
	if row == 1 && column == 1 {
		return 3
	}
	if row == 2 && column == 0 {
		return 4
	}
	if row == 2 && column == 1 {
		return 5
	}
	if row == 3 && column == 0 {
		return 6
	}
	return -1
}

func getNextPositionAndDirectionOnCube(current grid.Coords, direction int, p puzzle) (grid.Coords, int) {
	var next grid.Coords
	var nextDirection int
	// these are the lines just _inside_ the first, second, third, and fourth rows/columns
	s1, s2, s3, s4 := p.sideLength - 1, (2 * p.sideLength) - 1, (3 * p.sideLength) - 1, (4 * p.sideLength) - 1
	side := getSide(current, p)

	// hard-coded transitions between cube edges
	// only works for cube nets matching this format
	//  12
	//  3
	// 45
	// 6
	switch side {
	case 1:
		if direction == 3 {
			nextDirection = 0
			next = grid.Coords{ X: 0, Y: s4 - (s2 - current.X)}
		} else if direction == 2 {
			nextDirection = 0
			next = grid.Coords{ X: 0, Y: s3 - current.Y }
		}
	case 2:
		if direction == 3 {
			nextDirection = 3
			next = grid.Coords{ X: s1 - (s3 - current.X), Y: s4 }
		} else if direction == 0 {
			nextDirection = 2
			next = grid.Coords{ X: s2, Y: s3 - current.Y }
		} else if direction == 1 {
			nextDirection = 2
			next = grid.Coords{ X: s2, Y: s1 + (current.X - s2)}
		}
	case 3:
		if direction == 0 {
			nextDirection = 3
			next = grid.Coords{ X: s2 + (current.Y - s1), Y: s1 }
		} else if direction == 2 {
			nextDirection = 1
			next = grid.Coords{ X: s1 - (s2 - current.Y), Y: s2 + 1 }
		}
	case 4:
		if direction == 3 {
			nextDirection = 0
			next = grid.Coords{ X: s1 + 1, Y: s2 - (s1 - current.X) }
		} else if direction == 2 {
			nextDirection = 0
			next = grid.Coords{ X: s1 + 1, Y: s3 - current.Y}
		}
	case 5:
		if direction == 0 {
			nextDirection = 2
			next = grid.Coords{ X: s3, Y: s3 - current.Y}
		} else if direction == 1 {
			nextDirection = 2
			next = grid.Coords{ X: s1, Y: s3 + (current.X - s1)}
		}
	case 6:
		if direction == 0 {
			nextDirection = 3
			next = grid.Coords{ X: s1 + (current.Y - s3), Y: s3 }
		} else if direction == 1 {
			next = grid.Coords{ X: s3 - (s1 - current.X), Y: 0 }
			nextDirection = 1
		} else if direction == 2 {
			nextDirection = 1
			next = grid.Coords{ X: s2 - (s4 - current.Y), Y: 0 }
		}
	}
	return next, nextDirection
}