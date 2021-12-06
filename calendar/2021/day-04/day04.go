package main

import (
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

type space struct {
	value  string
	marked bool
}

func main() {
	input := files.ReadFile(04, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	drawnNumbers, boards := buildBoards(input)
	boardIndex := -1
	finalVal := -1
	for i, val := range drawnNumbers {
		boards = markBoards(val, boards)

		if i >= 4 {
			boardIndex = checkBoards(boards)
			if boardIndex >= 0 {
				finalVal, _ = strconv.Atoi(val)
				break
			}
		}
	}

	return finalVal * getUnmarkedSum(boards[boardIndex])
}

func solvePart2(input []string) int {
	drawnNumbers, boards := buildBoards(input)
	finalVal := -1
	lastBoard := [][]space{}
	for i, val := range drawnNumbers {
		boards = markBoards(val, boards)
		if i >= 4 {
			for {
				boardIndex := checkBoards(boards)
				if boardIndex >= 0 {
					finalVal, _ = strconv.Atoi(val)
					lastBoard = boards[boardIndex]
					ret := make([][][]space, 0)
					ret = append(ret, boards[:boardIndex]...)
					boards = append(ret, boards[boardIndex+1:]...)
				}
				if boardIndex < 0 {
					break
				}
			}

		}
	}

	return finalVal * getUnmarkedSum(lastBoard)
}

func buildBoards(input []string) ([]string, [][][]space) {
	drawnNumbers := strings.Split(input[0], ",")
	boards := [][][]space{}
	for i := 2; i < len(input); i += 6 {
		board := [][]space{}
		for j := i; j < i+5; j++ {
			row := []space{}
			splitString := strings.Split(input[j], " ")
			for _, val := range splitString {
				if val != "" {
					space := space{
						value:  val,
						marked: false,
					}
					row = append(row, space)
				}
			}
			board = append(board, row)
		}
		boards = append(boards, board)
	}
	return drawnNumbers, boards
}

func markBoards(val string, boards [][][]space) [][][]space {
	for _, board := range boards {
		marked := false
		for _, row := range board {
			for i := range row {
				if val == row[i].value {
					row[i].marked = true
					marked = true
					break
				}
				if marked {
					break
				}
			}
			if marked {
				break
			}
		}
	}
	return boards
}

func checkBoards(boards [][][]space) int {
	index := -1
	for ind, board := range boards {
		for i := 0; i < len(board); i++ {
			bingoCol, bingoRow := true, true
			for j := 0; j < len(board[i]); j++ {
				// check row
				if !board[j][i].marked {
					bingoCol = false
				}
				//check column
				if !board[i][j].marked {
					bingoRow = false
				}
				if !bingoCol && !bingoRow {
					break
				}
			}
			if bingoCol || bingoRow {
				index = ind
				break

			}
		}
	}
	return index
}

func getUnmarkedSum(board [][]space) int {
	sum := 0
	for _, row := range board {
		for _, space := range row {
			if !space.marked {
				valInt, _ := strconv.Atoi(space.value)
				sum += valInt
			}
		}
	}
	return sum
}

func printBoards(boards [][][]space) {
	for _, board := range boards {
		for _, row := range board {
			rowStrVals := []string{}
			for _, space := range row {
				if space.marked {
					rowStrVals = append(rowStrVals, space.value+"x")
				} else {
					rowStrVals = append(rowStrVals, space.value)

				}
			}
			println(strings.Join(rowStrVals, ","))
		}
		println("------------------")
	}
}
