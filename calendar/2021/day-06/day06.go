package main

import (
	"strconv"
	"strings"
	"tblue-aoc-2021/utils/files"
)

func main() {
	input := files.ReadFile(06, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	fishes := parseFishes(input[0])
	fishes = buildFishesForDays(fishes, 80)

	return len(fishes)
}

func solvePart2(input []string) int {
	fishes := parseFishesIntoMap(input[0])
	return trackFishesForDays(fishes, 256)

}

func parseFishes(input string) []int {
	fishesStrings := strings.Split(input, ",")
	fishes := []int{}

	for _, f := range fishesStrings {
		timer, _ := strconv.Atoi(f)
		fishes = append(fishes, timer)
	}
	return fishes
}

func buildFishesForDays(fishes []int, numDays int) []int {
	for i := 0; i < numDays; i++ {
		newFishes := 0
		for f := range fishes {
			if fishes[f] == 0 {
				newFishes++
				fishes[f] = 6
			} else {
				fishes[f] -= 1
			}
		}
		for j := 0; j < newFishes; j++ {
			fishes = append(fishes, 8)
		}
	}
	return fishes
}

func parseFishesIntoMap(input string) map[int]int {
	fishMap := make(map[int]int)
	fishesStrings := strings.Split(input, ",")

	for _, f := range fishesStrings {
		timer, _ := strconv.Atoi(f)
		fishMap[timer] = fishMap[timer] + 1
	}
	return fishMap
}
func trackFishesForDays(fishes map[int]int, numDays int) int {
	for i := 0; i < numDays; i++ {
		zeroFish := fishes[0]
		for ind := 1; ind <= 8; ind++ {
			fishes[ind-1] = fishes[ind]
		}
		fishes[6] += zeroFish
		fishes[8] = zeroFish
	}

	count := 0
	for _, value := range fishes {
		count += value
	}
	return count
}
