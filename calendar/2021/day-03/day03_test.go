package main

import (
	"testing"
)

func TestPart1SampleInput(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := int64(198)
	count := solvePart1(input)
	if count != want {
		t.Fatalf(`solvePart1(input) = %v, want match for %#v`, count, want)
	}
}

func TestPart2SampleInput(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	want := int64(230)
	count := solvePart2(input)
	if count != want {
		t.Fatalf(`solvePart2(input) = %v, want match for %#v`, count, want)
	}
}
