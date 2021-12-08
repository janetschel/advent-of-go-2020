package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	input := files.ReadFile(15, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	ingredients := parseInput(input)
	return buildRecipe(ingredients, 100, 0)
}

func solvePart2(input []string) int {
	ingredients := parseInput(input)
	return buildRecipe(ingredients, 100, 500)
}

// brute force, handles at most 4 ingredients
func buildRecipe(ingredients []ingredient, total int, calorieTarget int) int {
	maxScore := 0
	for i := 1; i <= total; i++ {
		for j := 1; j <= total; j++ {
			for k := 1; k <= total; k++ {
				l := total - i - j - k
				ratios := []int{i, j, k, l}[0:len(ingredients)]
				if slices.Sum(ratios) == total {
					score := scoreRecipe(ingredients, ratios)
					if score > maxScore && ((calorieTarget > 0 && countCalories(ingredients, ratios) == calorieTarget) || calorieTarget == 0) {
						maxScore = score
					}
				}
			}
		}
	}
	return maxScore
}

func scoreRecipe(ingredients []ingredient, amounts []int) int {
	capacity, durability, flavor, texture := 0, 0, 0, 0
	for i := range ingredients {
		amount := amounts[i]
		ingr := ingredients[i]
		capacity += amount * ingr.capacity
		durability += amount * ingr.durability
		flavor += amount * ingr.flavor
		texture += amount * ingr.texture
	}
	return slices.Max([]int{0, capacity}) * slices.Max([]int{0, durability}) * slices.Max([]int{0, flavor}) * slices.Max([]int{0, texture})
}

func countCalories(ingredients []ingredient, amounts []int) int {
	calories := 0
	for i := range ingredients {
		calories += amounts[i] * ingredients[i].calories
	}
	return calories
}

func parseInput(input []string) []ingredient {
	ingredients := []ingredient{}
	for i := range input {
		ingredients = append(ingredients, parseIngredient(input[i]))
	}
	return ingredients
}

func parseIngredient(input string) ingredient {
	parts := strings.Fields(input)
	capacity, _ := strconv.Atoi(parts[2][:len(parts[2])-1])
	durability, _ := strconv.Atoi(parts[4][:len(parts[4])-1])
	flavor, _ := strconv.Atoi(parts[6][:len(parts[6])-1])
	texture, _ := strconv.Atoi(parts[8][:len(parts[8])-1])
	calories, _ := strconv.Atoi(parts[10])
	return ingredient{
		name:       parts[0][:len(parts[0])-1],
		capacity:   capacity,
		durability: durability,
		flavor:     flavor,
		texture:    texture,
		calories:   calories,
	}
}
