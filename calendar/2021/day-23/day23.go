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
	s, halls := parseInput(input)
	return move(s, halls)
}

func solvePart2(input []string) int {
	// doesn't solve correctly for example ¯\_(ツ)_/¯
	newInput := make([]string, len(input))
	copy(newInput, input)
	newInput = append(newInput[0:5], newInput[3:]... )
	newInput[3], newInput[4] = "  #D#C#B#A#", "  #D#B#A#C#"
	s, halls := parseInput(newInput)
	return move(s, halls)
}

// this code isn't very readible, concise, or performant but it works, kind of
func move(s state, hallEntrances []int) int {
	sMem := s.memoized()
	stateLookup := map[string]*state{sMem: &s}
	minimumPrices := map[string]int{sMem: 0}
	queue := sets.New()
	queue.Add(sMem)
	winPrices := []int{}

	queueLoop:
		for queue.Size() > 0 {
			currentStateMemoized := queue.Random()
			queue.Remove(currentStateMemoized)
			currentState := stateLookup[currentStateMemoized]
			currentPrice := minimumPrices[currentStateMemoized]

			if currentState.isFinalState() {
				winPrices = append(winPrices, currentPrice)
				continue queueLoop
			}

			// move incorrect amphipods out of rooms into hallways
			moveFromRooms:
				for roomID := range currentState.rooms {
					extraPrice := 0
					itemsInRoom := 0

					allItemsMatch := true
					targetItem := roomOrder[roomID]
					for _, o := range currentState.rooms[roomID] {
						if o != "" {
							itemsInRoom++
						}
						if o != targetItem && o != "" {
							allItemsMatch = false
						}
					}
					if itemsInRoom == 0 || allItemsMatch {
						// if the room is empty, don't try to move things from it
						// if the room has only correct items don't move them out
						continue moveFromRooms
					}
					// add one to move out of entrance, direction is irrelevant
					extraPrice += len(currentState.rooms[roomID]) - itemsInRoom + 1

					beforeMovePrice := extraPrice
					possibleNewPlacesAndPrice := map[int]int{}
					pos := hallEntrances[roomID] - 1
					moveLeft:
						for pos >= 0 {
							if currentState.hallway[pos] == "" {
								extraPrice++
							} else {
								// collision, movement stops
								break moveLeft
							}
							possibleNewPlacesAndPrice[pos] = extraPrice
							pos--
						}
					extraPrice = beforeMovePrice
					pos = hallEntrances[roomID] + 1
					moveRight:
						for pos < len(currentState.hallway) {
							if currentState.hallway[pos] == "" {
								extraPrice++
							} else {
								// collision, movement stops
								break moveRight
							}
							possibleNewPlacesAndPrice[pos] = extraPrice
							pos++
						}
					
					validNewPlaces := map[int]int{}
					for position, price := range possibleNewPlacesAndPrice {
						// filter out hallway entrances as amphipods can't stop there
						if !slices.Contains(hallEntrances, position) {
							validNewPlaces[position] = price
						}
					}

					moved := ""
					movedFrom := 0
					newRooms:
						for i, o := range currentState.rooms[roomID] {
							if o != "" {
								moved = o
								movedFrom = i
								break newRooms
							}
						}

					for position, price := range validNewPlaces {
						newState := currentState.copy()
						
						// make the adjustments on the new state
						newState.rooms[roomID][movedFrom] = ""
						newState.hallway[position] = moved
						newPrice := currentPrice + (price * costLookup[moved])

						// todo: move this push/update to function
						mem := newState.memoized()
						previousPrice, knownPrice := minimumPrices[mem]
						if !knownPrice {
							previousPrice = math.MaxInt
						}

						if newPrice < previousPrice {
							queue.Add(mem)
							minimumPrices[mem] = newPrice
							stateLookup[mem] = &newState
						}
					}
				}

			// move out of halls into target rooms
			moveFromHalls:
				for hallLocation := range currentState.hallway {
					animalType := currentState.hallway[hallLocation]
					if animalType != "" {
						targetRoom := slices.IndexOf(animalType, roomOrder)
						howManyInFinalRoom := 0
						for _, o := range currentState.rooms[targetRoom] {
							if o != "" {
								howManyInFinalRoom++
							}
						}
						if howManyInFinalRoom == len(currentState.rooms[targetRoom]) {
							continue moveFromHalls
						} else if howManyInFinalRoom > 0 {
							differentTypes := false
							checkRoom:
								for i := 0; i < howManyInFinalRoom; i++ {
									if currentState.rooms[targetRoom][i] != animalType && currentState.rooms[targetRoom][i] != "" {
										differentTypes = true
										break checkRoom
									}
								}
							if differentTypes {
								continue moveFromHalls
							}
						}
						targetHallway := hallEntrances[targetRoom]
						step := -1
						if targetHallway > hallLocation {
							step = 1
						}
						position, collision, extraCost := hallLocation, false, 0
						for position != targetHallway {
							position += step
							if currentState.hallway[position] != "" {
								collision = true
							}
							extraCost += costLookup[animalType]
						}
						if collision {
							continue moveFromHalls
						}
						extraCost += (len(currentState.rooms[targetRoom]) - howManyInFinalRoom) * costLookup[animalType]

						newState := currentState.copy()
						newState.rooms[targetRoom][(len(currentState.rooms[targetRoom])-howManyInFinalRoom)-1] = animalType
						newState.hallway[hallLocation] = ""
						newPrice := currentPrice + extraCost
						
						mem := newState.memoized()
						previousPrice, knownPrice := minimumPrices[mem]
						if !knownPrice {
							previousPrice = math.MaxInt
						}

						if newPrice < previousPrice {
							queue.Add(mem)
							minimumPrices[mem] = newPrice
							stateLookup[mem] = &newState
						}
					}
				}
		}

	sort.Ints(winPrices)
	if len(winPrices) == 0 {
		return 0
	}
	return winPrices[0]
}

func isAmphipod(char string) bool {
	return char == amber || char == bronze || char == copper || char == desert
}

func parseInput(input []string) (state, []int) {
	hallway := make([]string, strings.Count(input[1], "."))
	rooms := make([][]string, 4)
	hallEntrances := []int{}
	for i := 0; i < len(input[2]); i++ {
		if isAmphipod(input[2][i:i+1]) {
			hallEntrances = append(hallEntrances, i - 1)
		}
	}

	for i := 2; i < len(input) - 1; i++ {
		roomValues := slices.Filter(strings.Split(input[i], ""), isAmphipod)
		for j, char := range roomValues {
			rooms[j] = append(rooms[j], char)
		}
	}

	return state{hallway: hallway, rooms: rooms }, hallEntrances
}

func (s state) isFinalState() bool {
	if strings.Trim(strings.Join(s.hallway, ""), " ") != "" {
		return false
	}
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
	return fmt.Sprintf("%v %v", s.hallway, s.rooms)
}

func (s state) copy() state {
	hallway := make([]string, len(s.hallway))
	copy(hallway, s.hallway)

	rooms := make([][]string, len(s.rooms))
	for i, room := range s.rooms {
		rooms[i] = make([]string, len(room))
		copy(rooms[i], s.rooms[i])
	}

	return state{hallway: hallway, rooms: rooms}
}
