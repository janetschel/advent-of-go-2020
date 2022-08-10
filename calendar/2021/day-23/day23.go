package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"fmt"
	"math"
	"sort"
	"strings"
)

type state struct {
	hallway       []string
	rooms         [][]string
	hallEntrances []int
}

var amber, bronze, copper, desert string = "A", "B", "C", "D"
var roomOrder []string = []string{amber, bronze, copper, desert}
var costLookup map[string]int = map[string]int{amber: 1, bronze: 10, copper: 100, desert: 1000}

func main() {
	input := files.ReadFile(23, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	// 18356 too high
	// next guess 18328

	s := parseInput(input)
	fmt.Println(move(s))

	return result
}

func solvePart2(input []string) int {
	result := 0

	return result
}

func move(s state) int {
	stateLookup := map[string]state{s.memoized(): s}
	queue := sets.New()
	queue.Add(s.memoized())
	bestPrice := map[string]int{s.memoized(): 0}
	winPrices := []int{}

	killswitch := 0
	for queue.Size() > 0 && killswitch < 1000000 {
		currentStateMemoized := queue.Iterator()[0]
		queue.Remove(currentStateMemoized)
		currentState := stateLookup[currentStateMemoized]
		currentPrice := bestPrice[currentStateMemoized]

		if currentState.isFinalState() {
			fmt.Println("found final state")
			winPrices = append(winPrices, currentPrice)
			continue
		}

		// move out of halls into target rooms
		for hallLoc := range currentState.hallway {
			animalType := currentState.hallway[hallLoc]
			if animalType != "" {
				targetRoom := slices.IndexOf(animalType, roomOrder)
				howManyInFinalRoom := 0
				for _, o := range currentState.rooms[targetRoom] {
					if o != "" {
						howManyInFinalRoom++
					}
				}
				if howManyInFinalRoom == len(currentState.rooms[targetRoom]) {
					continue
				} else if howManyInFinalRoom > 0 {
					differentTypes := false
					for i := 0; i < howManyInFinalRoom; i++ {
						if currentState.rooms[targetRoom][i] != animalType {
							differentTypes = true
							break
						}
					}
					if differentTypes {
						continue
					}
				}
				targetHallway := currentState.hallEntrances[targetRoom]
				step := -1
				if targetHallway > hallLoc {
					step = 1
				}
				position, collision, extraCost := hallLoc, false, 0
				for position != targetHallway {
					position += step
					if currentState.hallway[position] != "" {
						collision = true
					}
					extraCost += costLookup[animalType]
				}
				if collision {
					continue
				}
				extraCost += (len(currentState.rooms[targetRoom]) - howManyInFinalRoom) * costLookup[animalType]

				newState := currentState.copy()
				newState.rooms[targetRoom][(len(currentState.rooms[targetRoom])-howManyInFinalRoom)-1] = animalType
				newState.hallway[hallLoc] = ""
				newPrice := currentPrice + extraCost
				oldPriceOfNewState, inPrices := bestPrice[newState.memoized()]
				if howManyInFinalRoom == 1 {
					fmt.Printf("moving %v into room %v, room now has %v\n", animalType, targetRoom, newState.rooms[targetRoom])
				}
				if !inPrices {
					oldPriceOfNewState = math.MaxInt
				}
				if newPrice < oldPriceOfNewState {
					bestPrice[newState.memoized()] = newPrice
					queue.Add(newState.memoized())
					stateLookup[newState.memoized()] = newState
				}
			}
		}

		for roomID := range currentState.rooms {
			extraPrice := 0
			itemsInRoom := 0
			for _, o := range currentState.rooms[roomID] {
				if o != "" {
					itemsInRoom++
				}
			}
			if itemsInRoom == 0 {
				continue
			}
			allItemsMatch := true
			targetItem := roomOrder[roomID]
			for _, item := range currentState.rooms[roomID] {
				if item != targetItem {
					allItemsMatch = false
					break
				}
			}
			if allItemsMatch {
				continue
			}
			extraPrice += len(currentState.rooms[roomID]) - itemsInRoom
			extraPrice++

			beforeMovePrice := extraPrice
			possibleNewPlacesAndPrice := map[int]int{}
			pos := currentState.hallEntrances[roomID] - 1
			for pos >= 0 {
				if currentState.hallway[pos] == "" {
					extraPrice++
				} else {
					break
				}
				possibleNewPlacesAndPrice[pos] = extraPrice
				pos--
			}
			extraPrice = beforeMovePrice
			pos = currentState.hallEntrances[roomID] + 1
			for pos < len(currentState.hallway) {
				if currentState.hallway[pos] == "" {
					extraPrice++
				} else {
					break
				}
				possibleNewPlacesAndPrice[pos] = extraPrice
				pos++
			}

			validNewPlaces := map[int]int{}
			for position, price := range possibleNewPlacesAndPrice {
				isInEntrance := false
				for _, entrance := range currentState.hallEntrances {
					if position == entrance {
						isInEntrance = true
						break
					}
				}
				if !isInEntrance {
					validNewPlaces[position] = price
				}
			}
			for position, price := range validNewPlaces {
				newState := currentState.copy()
				// moved := newState.rooms[roomID][0]
				moved := ""
				movedFrom := 0
				for i, o := range newState.rooms[roomID] {
					if o != "" {
						moved = o
						movedFrom = i
						break
					}
				}
				newState.rooms[roomID][movedFrom] = ""
				newState.hallway[position] = moved
				newPrice := currentPrice + (price * costLookup[moved])
				oldPriceOfNewState, inMap := bestPrice[newState.memoized()]
				if !inMap {
					oldPriceOfNewState = math.MaxInt
				}
				if newPrice < oldPriceOfNewState {
					// fmt.Printf("Cost of moving %v from room %v to hallway pos %v is %v\n", moved, roomID, position, newPrice)
					bestPrice[newState.memoized()] = newPrice
					queue.Add(newState.memoized())
					stateLookup[newState.memoized()] = newState
				}
			}
		}

		killswitch++
	}

	fmt.Println(killswitch)
	// fmt.Println(moveCost)
	fmt.Println(winPrices)
	// fmt.Println(bestPrice)

	sort.Ints(winPrices)
	if len(winPrices) == 0 {
		return 0
	}
	return winPrices[0]
}

func parseInput(input []string) state {
	hallway := make([]string, strings.Count(input[1], "."))
	rooms := make([][]string, 4)
	hallEntrances := []int{}
	for i := range rooms {
		rooms[i] = make([]string, len(input)-3)
		roomIndex := (2 * (i + 1)) + 1
		hallEntrances = append(hallEntrances, roomIndex-1)
		for j := 2; j < len(input)-1; j++ {
			rooms[i][j-2] = input[j][roomIndex : roomIndex+1]
		}
	}

	return state{hallway: hallway, rooms: rooms, hallEntrances: hallEntrances}
}

func (s state) isFinalState() bool {
	for i, room := range s.rooms {
		for _, amphipod := range room {
			if amphipod != roomOrder[i] {
				return false
			}
		}
	}
	return true
}

func (s state) memoized() string {
	return fmt.Sprintf("%v", s)
}

func (s state) copy() state {
	hallway := make([]string, len(s.hallway))
	for i := range hallway {
		hallway[i] = s.hallway[i]
	}

	hallEntrances := make([]int, len(s.hallEntrances))
	for i := range hallEntrances {
		hallEntrances[i] = s.hallEntrances[i]
	}

	rooms := make([][]string, len(s.rooms))
	for i := range rooms {
		rooms[i] = make([]string, len(s.rooms[i]))
		for j := range rooms[i] {
			rooms[i][j] = s.rooms[i][j]
		}
	}

	return state{hallway: hallway, hallEntrances: hallEntrances, rooms: rooms}
}

/*
	currentState := queue[0]
	queue = queue[1:]

	currentCost := moveCost[currentState.memoized()]
	if currentState.isFinalState() {
		finalCosts = append(finalCosts, currentCost)
		fmt.Println("found valid final state")
		continue
	}

	for h := range currentState.hallway {
		occupant := currentState.hallway[h]
		if occupant != "" {
			costToMove := costLookup[occupant]
			targetRoom := slices.IndexOf(occupant, roomOrder)
			occupantCount, roomHasWrongAmp := 0, false
			for _, space := range currentState.rooms[targetRoom] {
				if space != "" {
					occupantCount++
				}
				if space != occupant {
					roomHasWrongAmp = true
				}
			}
			if occupantCount >= len(currentState.rooms[targetRoom]) || roomHasWrongAmp {
				continue
			}
			// if we get this far, the room has space and doesn't need to be emptied
			targetHallway := currentState.hallEntrances[targetRoom]
			direction := 1
			if targetHallway < h {
				direction = -1
			}
			currentPosition, collision, cost := h, false, 0
			for currentPosition != targetHallway {
				currentPosition += direction
				if currentState.hallway[currentPosition] != "" {
					collision = true
				}
				cost += costToMove
			}
			if collision {
				continue
			}
			if occupantCount == 0 {
				cost += 2 * costToMove
			} else {
				cost += costToMove
			}
			nextState := currentState.copy()
			nextState.rooms[targetRoom] = append([]string{occupant}, nextState.rooms[targetRoom]...)
			nextState.rooms[targetRoom] = nextState.rooms[targetRoom][0 : len(nextState.rooms[targetRoom])-1]
			nextState.hallway[h] = ""
			newCost := cost + currentCost
			oldCostOfNewState := moveCost[nextState.memoized()]
			if newCost < oldCostOfNewState || oldCostOfNewState == 0 {
				moveCost[nextState.memoized()] = newCost
				queue = append(queue, nextState)
			}
		}
	}

	for room := range currentState.rooms {
		additionalCost := 0
		occupants := []string{}
		for _, o := range currentState.rooms[room] {
			if o != "" {
				occupants = append(occupants, o)
			}
		}
		// fmt.Printf("working on occupants %v of room %v\n", occupants, room)
		if len(occupants) == 0 {
			// fmt.Printf("room %v empty\n", room)
			continue
		}
		if len(occupants) <= len(currentState.rooms[room]) {
			matchingCount := 0
			for _, o := range occupants {
				targetRoom := slices.IndexOf(o, roomOrder)
				if room == targetRoom {
					matchingCount++
				}
			}
			if matchingCount == len(currentState.rooms[room]) {
				fmt.Printf("room %v complete with %v of %v occupants\n", room, matchingCount, len(occupants))
				fmt.Println(currentState.memoized())
				continue
			}
			additionalCost += len(currentState.rooms[room]) - len(occupants) + 1
		}
		preMoveCost := additionalCost
		possibleMoves := map[int]int{}
		position := currentState.hallEntrances[room] - 1
		for position >= 0 {
			if currentState.hallway[position] == "" {
				additionalCost++
			} else {
				break
			}
			// fmt.Printf("Considering moving %v from room %v to hallway position %v\n", currentState.rooms[room][0], room, position)
			possibleMoves[position] = additionalCost
			position--
		}
		additionalCost = preMoveCost
		position = currentState.hallEntrances[room] + 1
		for position < len(currentState.hallway) {
			if currentState.hallway[position] == "" {
				additionalCost++
			} else {
				break
			}
			// fmt.Printf("Considering moving %v from room %v to hallway position %v\n", currentState.rooms[room][0], room, position)
			possibleMoves[position] = additionalCost
			position++
		}

		for move, cost := range possibleMoves {
			validMove := true
			for _, entrance := range currentState.hallEntrances {
				if entrance == move {
					validMove = false
					break
				}
			}
			if validMove {
				// fmt.Printf("Moving %v from room %v to hallway position %v is valid\n", currentState.rooms[room][0], room, move)
				newState := currentState.copy()
				moved := newState.rooms[room][0]
				newState.rooms[room][0] = ""
				newState.hallway[move] = moved
				newCost := currentCost + (cost * costLookup[moved])
				previousCost := moveCost[newState.memoized()]
				if newCost < previousCost || previousCost == 0 {
					moveCost[newState.memoized()] = newCost
					queue = append(queue, newState)
				}
				// fmt.Printf("move costs %v, previous cost %v\n", newCost, previousCost)
			} else {
				// fmt.Printf("Moving %v from room %v to hallway position %v is invalid\n", currentState.rooms[room][0], room, move)
			}
		}
	}

	fmt.Println(queue)
*/
