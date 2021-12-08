package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(06, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	fish := parseInput(input[0])
	result := simulate(fish, 80)
	return countFish(result)
}

func solvePart2(input []string) int {
	fish := parseInput(input[0])
	result := simulate(fish, 256)
	return countFish(result)
}

func simulate(fish []int, rounds int) map[int]int {
	fishMap := map[int]int{}
	for i := range fish {
		fishMap[fish[i]]++
	}
	for round := 0; round < rounds; round++ {
		roundMap := map[int]int{}
		for timer, count := range fishMap {
			if timer == 0 {
				roundMap[6] += count
				roundMap[8] += count
			} else {
				roundMap[timer-1] += count
			}
		}
		fishMap = roundMap
	}
	return fishMap
}

func countFish(fishMap map[int]int) int {
	count := 0
	for _, fishCount := range fishMap {
		count += fishCount
	}
	return count
}

func parseInput(input string) []int {
	fish := []int{}
	fishInput := strings.Split(input, ",")
	for i := range fishInput {
		val, _ := strconv.Atoi(fishInput[i])
		fish = append(fish, val)
	}
	return fish
}
