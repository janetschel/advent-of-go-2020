package main

import (
	"advent-of-go/utils/files"
	"sort"
	"strconv"
	"strings"
)

type chemical struct {
	quantity int
	name string
}

var ore, fuel = "ORE", "FUEL"
func main() {
	input := files.ReadFile(14, 2019, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	reactions := parseInput(input)
	return oreCostToProduceFuel(1, reactions)
}

func solvePart2(input []string) int {
	trillion := 1000000000000
	reactions := parseInput(input)
	return sort.Search(trillion, func(n int) bool {
		return oreCostToProduceFuel(n, reactions) > trillion
	}) - 1
}

func oreCostToProduceFuel(quantity int, reactions map[string][]chemical) int {
	return produceChemical(chemical{ name: fuel, quantity: quantity }, reactions, map[string]int{})
}

func produceChemical(chem chemical, reactions map[string][]chemical, stock map[string]int) int {
	name, quantity := chem.name, chem.quantity
	if name == ore {
		return quantity
	}
	
	if stock[name] >= quantity {
		stock[name] -= quantity
		return 0
	} else if stock[name] > 0 {
		inStock := stock[name]
		stock[name] = 0
		return produceChemical(chemical{ name: name, quantity: quantity - inStock }, reactions, stock)
	}

	reaction := reactions[name]
	needed, produced := reaction[:len(reaction) - 1], reaction[len(reaction) - 1]
	n := (quantity - 1) / produced.quantity + 1

	ores := 0
	for _, ing := range needed {
		ores += produceChemical(chemical{ name: ing.name, quantity: ing.quantity * n }, reactions, stock)
	}

	if (n * produced.quantity) - quantity > 0 {
		stock[name] += (n * produced.quantity) - quantity
	}
	return ores
}

func parseInput(input []string) map[string][]chemical {
	reactions := map[string][]chemical{}
	for _, line := range input {
		parts := strings.Split(line, " => ")
		inputs, output := parseChemicals(parts[0]), parseChemicals(parts[1])[0]
		// only one chemical is ever produced, used as the key and last element in value
		_, has := reactions[output.name]
		if has {
			println("already have reaction producing", output.name)
		}
		reactions[output.name] = append(inputs, output)
	}
	return reactions
}

func parseChemicals(chem string) []chemical {
	components := strings.Split(chem, ", ")
	chemicals := make([]chemical, len(components))
	for i, c := range components {
		parts := strings.Fields(c)
		quantity, _ := strconv.Atoi(parts[0])
		chemicals[i] = chemical{ quantity: quantity, name: parts[1] }
	}
	return chemicals
}
