package main

import (
	"advent-of-go-2020/utils/bins"
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	maps "advent-of-go-2020/utils/map"
	"advent-of-go-2020/utils/str"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(14, "\n")
	println("Solution part 1:", solvePart1(input))
}

func solvePart1(input []string) int {
	mem, currentMask := make(map[int64]int), ""

	for _, line := range input {
		if line[0:4] == "mask" {
			currentMask = strings.Replace(line, "mask = ", "", -1)
			continue
		}

		parts := strings.Split(line, " = ")

		addr := strings.Replace(strings.Split(parts[0], "mem[")[1], "]", "", -1)
		binary := bins.Pad(strconv.FormatInt(int64(conv.ToInt(parts[1])), 2), 36)

		for i := 0; i < len(currentMask); i++ {
			if string(currentMask[i]) != "X" {
				binary = str.ReplaceCharAt(binary, string(currentMask[i]), i)
			}
		}

		num, _ := strconv.ParseInt(binary, 2, 64)
		mem[int64(conv.ToInt(addr))] = int(num)
	}

    return maps.Sum(mem)
}
