package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
)

func main() {
	input := files.ReadFile(12, "\n")
	println(solvePart2(input))
}

func solvePart2(input []string) int {
	x, y, facing := 0,0, "E"
	xW, yW := 10, -1

	for _, element := range input {
		instruction := element[:1]
		amount := conv.ToInt(element[1:])

		x, y, xW, yW, facing = moveWaypoint(instruction, amount, facing, x, y, xW, yW)
		println("-----------------")
		println(x)
		println(y)
		println(xW)
		println(yW)
	}

	if x < 0 {
		x *= -1
	}

	if y < 0 {
		y *= - 1
	}

	return x+y
}

func moveWaypoint(instruction string, amount int, facing string, x int, y int, xW int, yW int) (int, int, int, int, string) {
	if instruction == "F" {
		x += xW * amount
		y += yW * amount
	} else if instruction == "N" {
		yW -= amount
	} else if instruction == "S" {
		yW += amount
	} else if instruction == "E" {
		xW += amount
	} else if instruction == "W" {
		xW -= amount
	} else if instruction == "R" {
		times := (amount / 90) % 4

		if times == 1 {
			oldX := xW
			xW = -yW
			yW = oldX
		} else if times == 2 {
			xW = -xW
			yW = -yW
		} else if times == 3 {
			oldX := xW
			xW = yW
			yW = -oldX
		}

	} else if instruction == "L" {
		times := ((-amount / 90) + 4) % 4
		if times == 1 {
			oldX := xW
			xW = -yW
			yW = oldX
		} else if times == 2 {
			xW = -xW
			yW = -yW
		} else if times == 3 {
			oldX := xW
			xW = yW
			yW = -oldX
		}

	}

	return x,y, xW, yW, facing
}
