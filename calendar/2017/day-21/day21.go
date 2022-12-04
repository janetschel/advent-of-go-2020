package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/grid"
	"strings"
)

func main() {
	input := files.ReadFile(21, 2017, "\n")
	println(generateArt(input, 5))
	println(generateArt(input, 18))
}

func generateArt(input []string, iterations int) int {
	startingPattern := [][]string { { ".", "#", "." }, { ".", ".", "#"}, { "#", "#", "#"}}
	rules := parseInput(input)
	for i := 0; i < iterations; i++ {
		startingPattern = transform(startingPattern, rules)
	}

	return countOnPixels(startingPattern)
}

func transform(pixels [][]string, rules map[string]string) [][]string {
	var chunkSize int
	x, y := grid.Size(pixels)
	if x % 2 == 0 && y % 2 == 0 {
		chunkSize = 2
	} else	if x % 3 == 0 && y % 3 == 0 {
		chunkSize = 3
	} else {
		panic("pixels size not divisible by two or three")
	}
	patterns := [][]string{}
	for i := 0; i < len(pixels); i += chunkSize {
		patterns = append(patterns, make([]string, x / chunkSize))
		for j := 0; j < len(pixels[i]); j += chunkSize {
			chunk := [][]string{}
			for ii := i; ii < i + chunkSize; ii++ {
				chunk = append(chunk, pixels[ii][j:j+chunkSize])
			}
			pattern := applyRule(chunk, rules)
			patterns[i / chunkSize][j / chunkSize] = pattern
		}
	}
	return hydrate(patterns)
}

func hydrate(patterns [][]string) [][]string {
	patternSize := strings.Count(patterns[0][0], "/") + 1
	hydrated := make([][]string, len(patterns) * patternSize)

	for i := 0; i < len(patterns); i++ {
		for k := 0; k < patternSize; k++ {
			hydrated[(i * patternSize) + k] = make([]string, len(patterns[0]) * patternSize)
		}
		for j := 0; j < len(patterns[i]); j++ {
			chunk := fromPatternString(patterns[i][j])
			for rowCount, row := range chunk {
				for colCount, value := range row {
					hydrated[(i * patternSize) + rowCount][(j * patternSize) + colCount] = value
				}
			}
		}
	}

	return hydrated
}

func toPatternString(pixels [][]string) string {
	rowStrings := []string{}
	for _, row := range pixels {
		rowStrings = append(rowStrings, strings.Join(row, ""))
	}
	return strings.Join(rowStrings, "/")
}

func fromPatternString(pattern string) [][]string {
	pixels := [][]string{}
	rows := strings.Split(pattern, "/")
	for _, row := range rows {
		values := strings.Split(row, "")
		pixels = append(pixels, values)
	}
	return pixels
}

func countOnPixels(pixels [][]string) int {
	pattern := toPatternString(pixels)
	return strings.Count(pattern, "#")
}

func parseInput(input []string) map[string]string {
	patternMap := map[string]string{}

	for _, rule := range input {
		parts := strings.Split(rule, " => ")
		patternMap[parts[0]] = parts[1]
	}

	return patternMap
}

func getPermutations(pixels [][]string) [][][]string {
	flippedHorizontal, flippedVertical := grid.FlipHorizontal(pixels), grid.FlipVertical(pixels)
	return [][][]string{
		pixels,
		flippedHorizontal,
		flippedVertical,
		grid.Rotate90(pixels),
		grid.Rotate180(pixels),
		grid.Rotate270(pixels),
		grid.Rotate90(flippedVertical),
		grid.Rotate180(flippedVertical),
		grid.Rotate270(flippedVertical),
	}
}

func applyRule(pixels [][]string, rules map[string]string) string {
	permutations := getPermutations(pixels)
	for _, p := range permutations {
		pattern := toPatternString(p)
		newPattern, match := rules[pattern]
		if match {
			return newPattern
		}
	}
	return ""
}
