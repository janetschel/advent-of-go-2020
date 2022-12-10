package grid

import (
	"advent-of-go/utils/maths"
	"fmt"
	"math"
)

// Coords represents a coordinate in a two-dimensional grid
type Coords struct {
	X int
	Y int
}

// Origin is a constant representing the origin in a 2D grid
var Origin Coords = Coords{ X: 0, Y: 0 }

// FlipVertical flips a grid along the x-axis
func FlipVertical[T comparable](grid [][]T) [][]T {
	flipped := [][]T{}

	for i := len(grid) - 1; i >= 0; i-- {
		row := make([]T, len(grid[i]))
		copy(row, grid[i])
		flipped = append(flipped, row)
	}

	return flipped
}

// FlipHorizontal flips a grid along the y-axis
func FlipHorizontal[T comparable](grid [][]T) [][]T {
	flipped := [][]T{}

	for _, row := range grid {
		newRow := make([]T, len(row))
		for i, value := range row {
			newRow[len(row) - i - 1] = value
		}
		flipped = append(flipped, newRow)
	}

	return flipped
}

// ToString creates a string representation of the grid
func ToString[T comparable](grid [][]T) string {
	gridString := ""
	for _, row := range grid {
		for _, value := range row {
			gridString += fmt.Sprint(value)
		}
		gridString += "\n"
	}
	return gridString
}

// Size returns the x and y dimensions of the grid
func Size[T comparable](grid [][]T) (int, int) {
	return len(grid[0]), len(grid)
}

// Rotate90 rotates the grid by 90 degress, clockwise
func Rotate90[T comparable](grid [][]T) [][]T {
	x, y := Size(grid)
	rotated := [][]T{}

	for i := 0; i < y; i++ {
		newRow := []T{}
		for j := 0; j < x; j++ {
			newRow = append(newRow, grid[y - j - 1][i])
		}
		rotated = append(rotated, newRow)
	}

	return rotated
}

// Rotate180 rotates the grid by 180 degress, clockwise
func Rotate180[T comparable](grid [][]T) [][]T {
	return Rotate90(Rotate90(grid))
}

// Rotate270 rotates the grid by 270 degress, clockwise
func Rotate270[T comparable](grid [][]T) [][]T {
	return Rotate90(Rotate90(Rotate90(grid)))
}

// RotateCoordsCounterclockwise rotates a set of 2D coordinates around a pivot point (counter-clockwise) by a given number of degrees
func RotateCoordsCounterclockwise(coords Coords, pivot Coords, degrees float64) Coords {
	rads := maths.DegreesToRadians(degrees)
	cos, sin := math.Cos(rads), math.Sin(rads)
	dx, dy := float64(coords.X - pivot.X), float64(coords.Y - pivot.Y)
	x := (dx * cos) - (dy * sin) + float64(pivot.X)
	y := (dx * sin) + (dy * cos) + float64(pivot.Y)
	return Coords{ X: int(x), Y: int(y)}
}

// PerimeterSize returns an int of the number of points along the perimeter of the grid
func PerimeterSize[T comparable](grid [][]T) int {
	m, n := len(grid[0]), len(grid)
	return (2 * (m - 1)) + (2 * (n - 1))
}

// IsInGrid determines if a given point is inside the bounds of a given grid
func IsInGrid[T comparable](coords Coords, grid [][]T) bool {
	return coords.X >= 0 && coords.Y >= 0 && coords.Y < len(grid) && coords.X < len(grid[coords.Y])
}

// ToString creates a unique string to represent coordinates
func (c Coords) ToString() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

// ManhattanDistance returns the Mannattan distance between two coords (orthogonal)
func (c Coords) ManhattanDistance(to Coords) int {
	return maths.Abs(c.X - to.X) + maths.Abs(c.Y - to.Y)
}
