package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(24, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	foundValidModelNumber := false
	i := 99999999999999
	for i >= 11111111111111 && !foundValidModelNumber {
		modelNumber := fmt.Sprint(i)
		if strings.Contains(modelNumber, "0") {
			i--
			continue
		}
		foundValidModelNumber = runProgram(input, modelNumber)
		if foundValidModelNumber {
			return i
		}
		fmt.Println(i)
		i--
	}
	return 0
}

func solvePart2(input []string) int {
	foundValidModelNumber := false
	i := 11111111111111
	for i <= 99999999999999 && !foundValidModelNumber {
		modelNumber := fmt.Sprint(i)
		if strings.Contains(modelNumber, "0") {
			i++
			continue
		}
		foundValidModelNumber = runProgram(input, modelNumber)
		if foundValidModelNumber {
			return i
		}
		i++
	}
	return 0
}

func runProgram(instructions []string, modelNumber string) bool {
	modelNumberAsStr := strings.Split(modelNumber, "")
	input := []int{}
	for i := range modelNumberAsStr {
		val, _ := strconv.Atoi(modelNumberAsStr[i])
		input = append(input, val)
	}
	registers := map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
	currentInput := 0
	for _, inst := range instructions {
		parts := strings.Fields(inst)
		switch parts[0] {
		case "inp":
			registers[parts[1]] = input[currentInput]
			currentInput++
		case "add":
			registers[parts[1]] += getB(parts[2], registers)
		case "mul":
			registers[parts[1]] *= getB(parts[2], registers)
		case "div":
			b := getB(parts[2], registers)
			if b == 0 {
				return false
			}
			registers[parts[1]] /= b
		case "mod":
			registers[parts[1]] %= getB(parts[2], registers)
		case "eql":
			if registers[parts[1]] == getB(parts[2], registers) {
				registers[parts[1]] = 1
			} else {
				registers[parts[1]] = 0
			}
		}
	}
	return registers["z"] == 0
}

func getB(b string, registers map[string]int) int {
	bValue, isRegister := registers[b]
	if isRegister {
		return bValue
	}
	bValue, _ = strconv.Atoi(b)
	return bValue
}
