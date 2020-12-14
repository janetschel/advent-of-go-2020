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
	println("Solution part 2:", solvePart2(input))
}

func solvePart2(input []string) int {
	mem, currentMask := make(map[int64]int), ""

	for _, line := range input {
		if line[0:4] == "mask" {
			currentMask = strings.Replace(line, "mask = ", "", -1)
			continue
		}

		split := strings.Split(line, " = ")

		addr := strings.Replace(strings.Split(split[0], "mem[")[1], "]", "", -1)
		addr = bins.Pad(strconv.FormatInt(int64(conv.ToInt(addr)), 2), 36)

		for i := 0; i < len(currentMask); i++ {
			if string(currentMask[i]) != "0" {
				addr = str.ReplaceCharAt(addr, string(currentMask[i]), i)
			}
		}

		binaries := bins.AllBinaryNumbers(strings.Count(currentMask, "X"))
		for _, memA := range permutation(addr, binaries) {
			memAddr, _ := strconv.ParseInt(memA, 2, 64)
			mem[memAddr] = conv.ToInt(split[1])
		}
	}

	return maps.Sum(mem)
}

func permutation(mask string, binaries []string) []string {
	addresses := make([]string, 0)

	for _, binaryNumber := range binaries {
		current := mask

		for i, j := 0, 0; i < len(current); i++ {
			if string(current[i]) == "X" {
				current = str.ReplaceCharAt(current, string(binaryNumber[j]), i)
				j++
			}
		}

		addresses = append(addresses, current)
	}

	return addresses
}
