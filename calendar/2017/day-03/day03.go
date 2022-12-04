package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"math"
	"strconv"
)

/*
1:  0,0
4:  0,1
9:  1,-1
16: -1,2
25: 2,-2
36: -3,3
49: 3,-3

largest numbers squares
layer 0 1 1
layer 1 9 3
layer 2 25 5
layer 3 49 7
layer 4 81 9
*/

func main() {
	input := files.ReadFile(03, 2017, "\n")
	value, _ := strconv.Atoi(input[0])
	println(solvePart1(value))
	println(solvePart2(value))
}

func solvePart1(input int) int {
	closestSquareRoot := int(math.Ceil(math.Sqrt(float64(input))))
	if int(closestSquareRoot) % 2 == 0 {
		closestSquareRoot++
	}
	layer := (closestSquareRoot - 1) / 2
	remainder := (closestSquareRoot * closestSquareRoot) - input
	pos := [2]int{layer, -1 * layer}
	if remainder == 0 {
		return int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
	}
	side := (remainder - 1) / (closestSquareRoot - 1)
	switch side {
	case 0:
		// 24-25
		pos[0] -= remainder
	case 1:
		// 20-17
		pos[0] -= (closestSquareRoot - 1)
		pos[1] += (remainder - (side * (closestSquareRoot - 1)))
	case 2:
		// 16-13
		pos[1] += (closestSquareRoot - 1)
		pos[0] -= (closestSquareRoot - 1) - (remainder - (side * (closestSquareRoot - 1)))
	case 3:
		// 12-10
		pos[1] += (closestSquareRoot - 1) - (remainder - (side * (closestSquareRoot - 1)))
	}

	return int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}

func solvePart2(input int) int {
	result := 0

	values := make(map[string]int)
	x, y := 0, 0
	// direction: 1 right 2 up 3 left 4 down
	direction, increment := 1,1

	values[getKey(x, y)] = 1

	for current := 0; current <= input; {
		for i := 1; i <= increment; i++ {
			switch direction {
			case 1:
				x++
			case 2:
				y++
			case 3:
				x--
			case 4:
				y--
			}
			current = getNeighbors(x,y,values)
			values[getKey(x,y)] = current
			if current > input {
				return current
			}
		}
		if direction == 4 {
			direction = 1
		} else {
			direction++
		}
		if direction == 1 || direction == 3 {
			increment++
		}
	}


	return result
}

func getKey(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func getNeighbors(x int, y int, values map[string]int) int {
	return values[getKey(x+1, y-1)] + values[getKey(x+1, y)] + values[getKey(x+1, y+1)] + values[getKey(x-1, y-1)] + values[getKey(x-1,y)] + values[getKey(x-1,y+1)] + values[getKey(x,y-1)] + values[getKey(x,y+1)]
}
