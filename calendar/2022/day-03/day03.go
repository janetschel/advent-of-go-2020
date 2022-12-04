package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"strings"
)

func main() {
	input := files.ReadFile(3, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	total := 0

	for _, rucksack := range input {
		misplaced := findMisplacedItem(rucksack)
		total += getItemValue(rune(misplaced))
	}

	return total
}

func solvePart2(input []string) int {
	total := 0

	for i := 0; i < len(input) - 2; i += 3 {
		set1, set2, set3 := sets.New(), sets.New(), sets.New()
		set1.AddRange(strings.Split(input[i], ""))
		set2.AddRange(strings.Split(input[i+1], ""))
		set3.AddRange(strings.Split(input[i+2], ""))
		common := findCommonItem([]sets.Set{set1, set2, set3})
		total += getItemValue(rune(common))
	}

	return total
}

func findCommonItem(itemSets []sets.Set) byte {
	intersection := itemSets[0]
	for i := 1; i < len(itemSets); i++ {
		intersection = intersection.Intersect(itemSets[i])
	}
	return intersection.Iterator()[0][0]
}

func findMisplacedItem(rucksack string) byte {
	compartment1, compartment2 := sets.New(), sets.New()
	half := len(rucksack) / 2
	for i := 0; i < half; i++ {
		compartment1.Add(rucksack[i:i+1])
		compartment2.Add(rucksack[half + i:half+i+1])
	}
	return findCommonItem([]sets.Set{ compartment1, compartment2 })
}

func getItemValue(item rune) int {
	value := int(item)
	if value >= 65 && value <= 90 {
		return value - 38
	}
	return value - 96
}
