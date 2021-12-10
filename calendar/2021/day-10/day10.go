package main

import (
	"sort"
	"tblue-aoc-2021/utils/files"
)

type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str rune) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	input := files.ReadFile(10, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	r, _ := evaluateLines(input)
	return calculateScore(r)
}

func solvePart2(input []string) int {
	_, s := evaluateLines(input)
	return calculateAutocompleteScore(s)
}

func evaluateLines(input []string) ([]rune, []Stack) {
	invalidTokens := []rune{}
	remainingStacks := []Stack{}
	for _, s := range input {
		stack := Stack{}
		valid := true
		for _, r := range s {
			switch r {
			case '(', '[', '{', '<':
				stack.Push(r)
			case ')':
				or, _ := stack.Pop()
				if or != '(' {
					invalidTokens = append(invalidTokens, r)
					valid = false
				}
			case ']':
				or, _ := stack.Pop()
				if or != '[' {
					invalidTokens = append(invalidTokens, r)
					valid = false
				}
			case '}':
				or, _ := stack.Pop()
				if or != '{' {
					invalidTokens = append(invalidTokens, r)
					valid = false
				}
			case '>':
				or, _ := stack.Pop()
				if or != '<' {
					invalidTokens = append(invalidTokens, r)
					valid = false
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			remainingStacks = append(remainingStacks, stack)
		}
	}
	return invalidTokens, remainingStacks
}

func calculateScore(invalidTokens []rune) int {
	sum := 0
	for _, r := range invalidTokens {
		switch r {
		case ')':
			sum += 3
		case ']':
			sum += 57
		case '}':
			sum += 1197
		case '>':
			sum += 25137
		}
	}
	return sum
}

func calculateAutocompleteScore(stacks []Stack) int {
	scores := []int{}
	for _, s := range stacks {
		total := 0
		for !s.IsEmpty() {
			r, _ := s.Pop()
			value := 0
			switch r {
			case '(':
				value = 1
			case '[':
				value = 2
			case '{':
				value = 3
			case '<':
				value = 4
			}
			total = (total * 5) + value
		}
		scores = append(scores, total)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
