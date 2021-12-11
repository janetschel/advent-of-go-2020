package main

import (
	"strconv"
	"tblue-aoc-2021/utils/files"
)

type Stack [][]int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(pair []int) {
	*s = append(*s, pair) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() ([]int, bool) {
	if s.IsEmpty() {
		return []int{}, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	input := files.ReadFile(11, 2021, "\n", false)
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	octoGrid := parseInput(input)

	return countFlashesForDays(octoGrid, 100)
}

func solvePart2(input []string) int {
	octoGrid := parseInput(input)

	return calculateDayWhenAllFlash(octoGrid)
}

func parseInput(input []string) [][]int {
	octoGrid := [][]int{}
	for _, val := range input {
		row := []int{}
		for _, r := range val {
			xstr := string(r) //x is rune converted to string
			num, _ := strconv.Atoi(xstr)
			row = append(row, num)
		}
		octoGrid = append(octoGrid, row)
	}
	return octoGrid
}

func countFlashesForDays(octoGrid [][]int, days int) int {
	count := 0
	for i := 0; i < days; i++ {
		flashed := map[string]bool{}
		for x, row := range octoGrid {
			for y := range row {
				stack := Stack{[]int{x, y}}
				count += increaseFromStack(octoGrid, stack, flashed)
			}
		}
		//printGrid(octoGrid)
	}
	return count
}

func increaseFromStack(octoGrid [][]int, stack Stack, flashed map[string]bool) int {
	count := 0
	for !stack.IsEmpty() {
		pair, _ := stack.Pop()
		x := pair[0]
		y := pair[1]
		val := octoGrid[x][y]
		key := strconv.Itoa(x) + "," + strconv.Itoa(y)
		_, found := flashed[key]
		if val == 0 && found {
			//skip, and we will have already take care of it
			continue
		}
		if val == 9 {
			// set to zero, increase all of its neighbors
			octoGrid[x][y] = 0
			count++
			flashed[key] = true
			stack = addNeighborsToStack(octoGrid, stack, x, y)
			continue
		}
		octoGrid[x][y]++
	}
	return count
}
func addNeighborsToStack(octoGrid [][]int, stack Stack, x int, y int) Stack {
	if x-1 >= 0 {
		stack.Push([]int{x - 1, y})
	}
	if y-1 >= 0 {
		stack.Push([]int{x, y - 1})
	}
	if x+1 < len(octoGrid) {
		stack.Push([]int{x + 1, y})
	}
	if y+1 < len(octoGrid[x]) {
		stack.Push([]int{x, y + 1})
	}
	if x-1 >= 0 && y-1 >= 0 {
		stack.Push([]int{x - 1, y - 1})
	}
	if x+1 < len(octoGrid) && y-1 >= 0 {
		stack.Push([]int{x + 1, y - 1})
	}
	if x+1 < len(octoGrid) && y+1 < len(octoGrid[x]) {
		stack.Push([]int{x + 1, y + 1})
	}
	if x-1 >= 0 && y+1 < len(octoGrid[x]) {
		stack.Push([]int{x - 1, y + 1})
	}
	return stack
}

func calculateDayWhenAllFlash(octoGrid [][]int) int {
	count := 0
	day := 0
	for count != 100 {
		count = 0
		day++
		flashed := map[string]bool{}
		for x, row := range octoGrid {
			for y := range row {
				stack := Stack{[]int{x, y}}
				count += increaseFromStack(octoGrid, stack, flashed)
			}
		}
		//printGrid(octoGrid)
	}
	return day
}

// func printGrid(octoGrid [][]int) {
// 	for _, row := range octoGrid {
// 		s, _ := json.Marshal(row)
// 		println(strings.Trim(string(s), "[]"))
// 	}
// 	println()
// }
