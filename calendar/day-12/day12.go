package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
)

func main() {
	input := files.ReadFile(12, "\n")
	println(solvePart1(input))
}

func solvePart1(input []string) int {
    x, y, facing := 0,0, "E"

    for _, element := range input {
		instruction := element[:1]
		amount := conv.ToInt(element[1:])

		x, y, facing = move(instruction, amount, facing, x, y)
	}

	if x < 0 {
		x *= -1
	}

	if y < 0 {
		y *= - 1
	}

    return x+y
}

func move(instruction string, amount int, facing string, x int, y int) (int, int, string) {
	if instruction == "F" {
		if facing == "E" {
			x += amount
		} else if facing == "N" {
			y -= amount
		} else if facing == "S" {
			y += amount
		} else if facing == "W"{
			x -= amount
		}
	} else if instruction == "N" {
		y -= amount
	} else if instruction == "S" {
		y += amount
	} else if instruction == "E" {
		x += amount
	} else if instruction == "W" {
		x -= amount
	} else if instruction == "R" {
		times := amount % 89

		if facing == "E" {
			if times == 1 {
				facing = "S"
			} else if times == 2 {
				facing = "W"
			} else if times == 3 {
				facing = "N"
			}
		} else if facing == "S" {
			if times == 1 {
				facing = "W"
			} else if times == 2 {
				facing = "N"
			} else if times == 3 {
				facing = "E"
			}
		} else if facing == "N" {
			if times == 1 {
				facing = "E"
			} else if times == 2 {
				facing = "S"
			} else if times == 3 {
				facing = "W"
			}
		} else if facing == "W" {
			if times == 1 {
				facing = "N"
			} else if times == 2 {
				facing = "E"
			} else if times == 3 {
				facing = "S"
			}
		}
	} else if instruction == "L" {
		times := amount % 89
		if facing == "E" {
			if times == 1 {
				facing = "N"
			} else if times == 2 {
				facing = "W"
			} else if times == 3 {
				facing = "S"
			}
		} else if facing == "S" {
			if times == 1 {
				facing = "E"
			} else if times == 2 {
				facing = "N"
			} else if times == 3 {
				facing = "W"
			}
		} else if facing == "N" {
			if times == 1 {
				facing = "W"
			} else if times == 2 {
				facing = "S"
			} else if times == 3 {
				facing = "E"
			}
		} else if facing == "W" {
			if times == 1 {
				facing = "S"
			} else if times == 2 {
				facing = "E"
			} else if times == 3 {
				facing = "N"
			}
		}

	}

	return x,y,facing
}
