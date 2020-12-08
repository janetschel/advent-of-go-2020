package handheld_computer_test

import (
	"advent-of-go-2020/handheld-computer"
	"strings"
	"testing"
)

func TestCorrectACC(t *testing.T) {
	input := strings.Split(
		"acc +10\n" +
		"acc +10\n" +
		"nop -20\n" +
		"jmp +2\n" +
		"acc -100\n" +
		"acc +90\n" +
		"nop +10", "\n")

	computer := handheld_computer.CreateComputer(input)
	result := computer.Execute()

	validate(t, result, 110)
}

func TestCorrectNOP(t *testing.T) {
	input := strings.Split(
		 "acc +20\n" +
			"nop -10\n" +
			"nop -10\n" +
			"acc +0", "\n")

	computer := handheld_computer.CreateComputer(input)
	result := computer.Execute()

	validate(t, result, 20)
}

func TestCorrectJMP(t *testing.T) {
	input := strings.Split(
		 "jmp 3\n" +
			"acc +10\n" +
		 	"acc +10\n" +
		 	"nop +10", "\n")

	computer := handheld_computer.CreateComputer(input)
	result := computer.Execute()

	validate(t, result, 0)
}

func TestCorrectACCAndJMP(t *testing.T) {
	input := strings.Split(
		 "acc +10\n" +
			"acc +10\n" +
		 	"acc +10\n" +
		 	"acc +10\n" +
		 	"acc +10\n" +
		 	"acc +10\n" +
		 	"jmp 2\n" +
		 	"jmp -4\n" +
		 	"nop +1000", "\n")

	computer := handheld_computer.CreateComputer(input)
	result := computer.Execute()

	validate(t, result, 60)
}


func validate(t *testing.T, result int, expected int) {
	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
}
