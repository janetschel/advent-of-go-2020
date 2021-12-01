package main

import (
	"advent-of-go/utils/files"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(05, 2020, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) uint64 {
	var max uint64 = 0;

	for _, seat := range input {
		seatId := parseSeat(seat)
		if seatId > max {
			max = seatId
		}
	}

	return max
}

func solvePart2(input []string) uint64 {
	ids := make([]uint64, len(input))
	for i := range input {
		ids = append(ids, parseSeat(input[i]))
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	
	for i := range ids {
		if ids[i] + 1 != ids[i + 1] && ids[i] != 0 {
			return ids[i] + 1
		}
	}

	return 0
}

func parseSeat(seat string) uint64 {
	rowIdentifier := seat[:len(seat) - 3]
	columnIdentifier := seat[len(seat) - 3:]

	rowIdentifier = strings.ReplaceAll(rowIdentifier, "F", "0")
	rowIdentifier = strings.ReplaceAll(rowIdentifier, "B", "1")

	columnIdentifier = strings.ReplaceAll(columnIdentifier, "L", "0")
	columnIdentifier = strings.ReplaceAll(columnIdentifier, "R", "1")

	row, _ := strconv.ParseUint(rowIdentifier, 2, 64)
	column, _ := strconv.ParseUint(columnIdentifier, 2, 64)

	return getSeatId(row, column)
}

func getSeatId(row uint64, column uint64) uint64 {
	return (row * 8) + column
}
