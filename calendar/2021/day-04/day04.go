package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

type coordinates struct {
	x      int
	y      int
	marked bool
}

type board struct {
	// convenience slices for checking bingo without iterating over board
	columns  []int
	rows     []int
	boardMap map[int]*coordinates
	hasBingo bool
}

type bingoGame struct {
	numbers []int
	boards  []board
}

type gameResult struct {
	winningBoardIndex int
	lastNumberPlayed  int
}

func main() {
	input := files.ReadFile(04, 2021, "\n\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	game := parseBingoGame(input)
	result := game.playGame()[0]
	return game.boards[result.winningBoardIndex].scoreBoard(result.lastNumberPlayed)
}

func solvePart2(input []string) int {
	game := parseBingoGame(input)
	results := game.playGame()
	result := results[len(results)-1]
	return game.boards[result.winningBoardIndex].scoreBoard(result.lastNumberPlayed)
}

func parseBoard(input string) board {
	rows := strings.Split(input, "\n")
	boardMap := map[int]*coordinates{}
	for y := range rows {
		values := strings.Fields(rows[y])
		for x := range values {
			value, _ := strconv.Atoi(values[x])
			boardMap[value] = &coordinates{x: x, y: y, marked: false}
		}
	}

	return board{
		boardMap: boardMap,
		rows:     make([]int, len(rows)),
		columns:  make([]int, len(boardMap)/len(rows)),
		hasBingo: false,
	}
}

func parseBingoGame(input []string) bingoGame {
	numbers := []int{}
	numbersStr := strings.Split(input[0], ",")
	for i := range numbersStr {
		value, _ := strconv.Atoi(numbersStr[i])
		numbers = append(numbers, value)
	}

	boards := []board{}
	for _, b := range input[1:] {
		boards = append(boards, parseBoard(b))
	}

	return bingoGame{
		numbers: numbers,
		boards:  boards,
	}
}

func (b *board) markNumber(number int) {
	coords, present := b.boardMap[number]
	if present {
		coords.marked = true
		b.columns[coords.x]++
		b.rows[coords.y]++

		b.hasBingo = b.columns[coords.x] == len(b.rows) || b.rows[coords.y] == len(b.columns)
	}
}

func (game *bingoGame) playGame() []gameResult {
	results := []gameResult{}
	for i := range game.numbers {
		for j := range game.boards {
			b := &game.boards[j]
			if !b.hasBingo {
				b.markNumber(game.numbers[i])
				if b.hasBingo {
					results = append(results, gameResult{
						winningBoardIndex: j,
						lastNumberPlayed:  game.numbers[i],
					})
				}
			}
		}
	}
	return results
}

func (b *board) scoreBoard(lastNumberPlayed int) int {
	if !b.hasBingo {
		return -1
	}

	if lastNumberPlayed == 0 {
		return 0
	}

	sum := 0
	for value, coords := range b.boardMap {
		if !coords.marked {
			sum += value
		}
	}
	return sum * lastNumberPlayed
}
