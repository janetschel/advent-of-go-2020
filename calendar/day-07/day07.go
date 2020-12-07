package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"go/types"
	"regexp"
	"strings"
)

// Datastructure for problem
type Bag struct {
	color string
	contains []InnerBag
}

type InnerBag struct {
	color string
	times int
}

func (b *Bag) init(color string) {
	b.color = color
}

func (b *Bag) addInner(inner InnerBag) {
	b.contains = append(b.contains, inner)
}

func (i *InnerBag) init(color string, times int) {
	i.color = color
	i.times = times
}

// Solution begins here
// Solution for part 1 and 2 are in the same file since they need the same functions and structs, so splitting them up
// does not make very much sense
func main() {
	input := files.ReadFile(7, "\n")
	solutionPart1 := solve(input)
	solutionPart2 := solvePart2(input)

	println("Solution part 1:", solutionPart1)
	println("Solution part 2:", solutionPart2)
}

func solve(input []string) int {
	bags := make([]Bag, 0)

	for _, str := range input {
		bags = append(bags, createBagFromString(str))
	}

	unique := make(map[string]types.Nil)
	search(bags, "shiny gold", &unique)
	return len(unique)
}

func solvePart2(input []string) int {
	bags := make([]Bag, 0)

	for _, str := range input {
		bags = append(bags, createBagFromString(str))
	}

	return searchPart2(bags, "shiny gold")
}

func search(bags []Bag, targetColor string, known *map[string]types.Nil) {
	containing := bagsContainedIn(bags, targetColor)

	for key := range containing {
		search(bags, key, known)
		(*known)[key] = types.Nil{}
	}
}

func searchPart2(bags []Bag, targetColor string) int {
	result := 0

	for _, outer := range bags {
		for _, inner := range outer.contains {
			if outer.color == targetColor {
				result += inner.times + inner.times * searchPart2(bags, inner.color)
			}
		}
	}

	return result
}

func bagsContainedIn(bags []Bag, targetColor string) map[string]types.Nil {
	// used for linear lookup-time, just like a set
	containing := make(map[string]types.Nil)

	for _, outer := range bags {
		for _, inner := range outer.contains {
			if inner.color == targetColor {
				containing[outer.color] = types.Nil{}
			}
		}
	}

	return containing
}

func createBagFromString(input string) Bag {
	regex := "^((?:\\w+)? \\w+) bags contain (?:(no other bags.)|" +
		"((?:(?:(?:\\d+) (?:(?:\\w+)? \\w+) (?:bags|bag), )|(?:(?:\\d+) (?:(?:\\w+)? \\w+) (?:bags|bag)\\.))*))$"

	return *createBag(regexp.MustCompile(regex).FindStringSubmatch(input))
}

func createBag(matches []string) *Bag{
	outer := new(Bag)
	outer.init(matches[1])

	if matches[2] == "no other bags." {
		return outer
	}

	inners := strings.Split(strings.Replace(matches[3], ".", "", 1), ", ")
	for _, innerString := range inners {
		parts := regexp.MustCompile("^(\\d+) ((?:\\w)*(?: |)\\w+) (?:bags|bag)$").FindStringSubmatch(innerString)

		inner := new(InnerBag)
		inner.init(parts[2], conv.ToInt(parts[1]))
		outer.addInner(*inner)
	}

	return outer
}
