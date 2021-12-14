package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
)

type Bag struct {
	count int
	color string
	contents []Bag
}

func main() {
	input := files.ReadFile(07, 2020, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	bags := buildBagList(input)
	fmt.Printf("%v\n", bags)

	return result
}

func buildBagList(input []string) []Bag {
	bags := make([]Bag, 0)
	for i := range input {
		bags = append(bags, parseLine(input[i]))
	}
	return bags
}

func parseLine(line string) Bag {
	topLevelParts := strings.Split(line, " bags contain ")
	innerBags := strings.Split(topLevelParts[1], ", ")
	contents := make([]Bag, 0)
	for i := range innerBags {
		bag := strings.ReplaceAll(innerBags[i], "bags", "")
		bag = strings.ReplaceAll(bag, "bag", "")
		bag = strings.ReplaceAll(bag, ".", "")
		bagParts := strings.Split(bag, " ")
		count, _ := strconv.Atoi(bagParts[0])
		contents = append(contents, Bag{
			count: count,
			color: strings.Join(bagParts[1:], " "),
		})
	}
	return Bag {
		color: topLevelParts[0],
		count: 1,
		contents: contents,
	}
}

func searchForParentsOf(children []Bag, bags []Bag) []Bag {
	parents := make([]Bag, 0)
	childColors = children.map { $0.color }
	for i := range bags {
		matchingChildren := false
		j := 0
		while j < len(bags[i].contents) && !matchingChildren {
			j++
		}
		if matchingChildren {
			parents = append(parents, bags[i])
		}
	}
	if len(parents) == 0 return []
	return append(parents, searchForParentsOf(parents, bags))
}

func solvePart2(input []string) int {
	result := 0



	return result
}
