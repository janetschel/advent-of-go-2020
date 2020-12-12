package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/maths"
)

type Point struct {
	x, y int
}

func main() {
	input := files.ReadFile(12, "\n")
	println("Ships Manhattan distance from (0, 0):", solvePart1(input))
	println("Ships Manhattan distance from (0, 0) (pt. 2):", solvePart2(input))
}

func solvePart1(input []string) int {
	ship := Point{
		x: 0,
		y: 0,
	}

	facing := 0
	for _, element := range input {
		facing = executeInstruction(element[:1], conv.ToInt(element[1:]), facing, &ship)
	}

	return maths.Abs(ship.x) + maths.Abs(ship.y)
}

func solvePart2(input []string) int {
	ship := Point {
		x: 0,
		y: 0,
	}

	waypoint := Point {
		x: 10,
		y: -1,
	}

	for _, element := range input {
		executeWaypointInstruction(element[:1], conv.ToInt(element[1:]), &ship, &waypoint)
	}

	return maths.Abs(ship.x) + maths.Abs(ship.y)
}

func executeInstruction(instruction string, amount int, facing int, ship *Point) int {
	directions := []string {"E", "S", "W", "N"}

	if instruction == "N" || instruction == "S" || instruction == "E" || instruction == "W" {
		ship.move(amount, instruction)
	} else if instruction == "F" {
		ship.move(amount, directions[facing])
	} else if instruction == "R" {
		facing = (facing + (amount % 89)) % 4
	} else if instruction == "L" {
		facing = (facing - (amount % 89) + 4) % 4
	}

	return facing
}

func executeWaypointInstruction(instruction string, amount int, ship *Point, waypoint *Point) {
	if instruction == "N" || instruction == "S" || instruction == "E" || instruction == "W" {
		waypoint.move(amount, instruction)
	} else if instruction == "R" {
		waypoint.rotateAroundShip((amount % 89) % 4)
	} else if instruction == "L" {
		waypoint.rotateAroundShip(((-amount % 89) + 4) % 4)
	} else if instruction == "F" {
		ship.x += waypoint.x * amount
		ship.y += waypoint.y * amount
	}
}

func (point *Point) rotateAroundShip(amount int) {
	if amount == 1 {
		oldWaypointX := point.x
		point.x = -point.y
		point.y = oldWaypointX
	} else if amount == 2 {
		point.x = -point.x
		point.y = -point.y
	} else if amount == 3 {
		oldWaypointX := point.x
		point.x = point.y
		point.y = -oldWaypointX
	}
}

func (point *Point) move(amount int, direction string) {
	if direction == "N" {
		point.y -= amount
	} else if direction == "S" {
		point.y += amount
	} else if direction == "E" {
		point.x += amount
	} else if direction == "W" {
		point.x -= amount
	}
}
