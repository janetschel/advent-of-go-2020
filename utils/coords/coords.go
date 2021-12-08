package coords

import "reflect"

type Coordinates interface {
	getX() int
	getY() int
	initialize(x int, y int)
}

func GetNeighbors(coords Coordinates, gridSize Coordinates) []Coordinates {
	coordsType := reflect.TypeOf(coords).Elem()
	neighbors := []Coordinates {}
	for x := -1; x <= 1 && x + coords.getX() < gridSize.getY(); x++ {
		for y := -1; y <= 1 && y + coords.getY() < gridSize.getY(); y++ {
			newX, newY := x + coords.getX(), y + coords.getY()
			if newX >= 0 && newY >= 0 && !(newX == coords.getX() && newY == coords.getY()) {
				newCoords := reflect.New(coordsType).Interface().(Coordinates)
				newCoords.initialize(newX, newY)
				neighbors = append(neighbors, newCoords)
			}
		}
	}

	return neighbors
}
