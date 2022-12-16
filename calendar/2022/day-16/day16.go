package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"advent-of-go/utils/slices"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type valve struct {
	name string
	flowRate int
	tunnelsTo []string
}

type state struct {
	current string
	timeRemaining int
	pressure int
	open []string
	path string
}

func main() {
	input := files.ReadFile(16, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	valves := parseInput(input)

	starting := state{ current: "AA", timeRemaining: 30, pressure: 0, open: []string{} }
	lookup := map[string]*state{ starting.ToString(): &starting }
	queue := sets.New()
	queue.Add(starting.ToString())

	endingPressures := []int{}

	killswitch := 0
	for queue.Size() > 0 && killswitch < 9999999 {
		currString := queue.Random()
		currentState := lookup[currString]
		currentState.updateTimeAndPressure(valves)

		queue.Remove(currString)
		lookup[currString] = nil

		lookup[currentState.ToString()] = currentState

		if currentState.timeRemaining <= 0 {
			if currentState.pressure > 1600 {
				fmt.Println(currentState)
			}
			endingPressures = append(endingPressures, currentState.pressure)
		} else {
			generatedStates := 0
			// if current has a flow rate and is not open, open in
			if !slices.Contains(currentState.open, currentState.current) && valves[currentState.current].flowRate > 0 {
				newState := currentState.copy()
				newState.open = append(newState.open, currentState.current)
				lookup[newState.ToString()] = &newState
				queue.Add(newState.ToString())
				generatedStates++
			}

			for _, v := range valves[currentState.current].tunnelsTo {
				// visit the valves which haven't been opened
				// if !slices.Contains(currentState.open, v) {
				visits := strings.Count(currentState.path, v)
				if visits < len(valves[v].tunnelsTo) {
					newState := currentState.copy()
					newState.current = v
					newState.path += v + " "
					lookup[newState.ToString()] = &newState
					queue.Add(newState.ToString())
					generatedStates++
				}
			}

			// if no moves/opens are relevant, stay in place
			if generatedStates == 0 {
				for currentState.timeRemaining > 0 {
					currentState.updateTimeAndPressure(valves)
				}
				if currentState.pressure > 1600 {
					println("finishing")
					fmt.Println(currentState)
				}
				endingPressures = append(endingPressures, currentState.pressure)
				// queue.Add(currentState.ToString())
			}
		}
		killswitch++
	}

	println(killswitch, queue.Size())
	if len(endingPressures) == 0 {
		return -1
	}
	sort.Ints(endingPressures)
	return endingPressures[len(endingPressures) - 1]
}

func solvePart2(input []string) int {
	result := 0



	return result
}

func parseInput(input []string) (map[string]valve) {
	valves := map[string]valve{}
	for _, v := range input {
		parts := strings.FieldsFunc(v, func(r rune) (bool) {
			return strings.ContainsRune(" =;,", r)
		})
		flowRate, _ := strconv.Atoi(parts[5])
		valve := valve{
			name: parts[1],
			flowRate: flowRate,
			tunnelsTo: parts[10:],
		}
		valves[parts[1]] = valve
	}
	return valves
}

func (s state) ToString() string {
	return fmt.Sprintf("%s %d %d %v", s.current, s.timeRemaining, s.pressure, s.open)
}

func (s state) copy() state {
	newOpen := make([]string, len(s.open))
	copy(newOpen, s.open)
	return state{
		current: s.current,
		timeRemaining: s.timeRemaining,
		pressure: s.pressure,
		open: newOpen,
		path: s.path,
	}
}

func (s *state) updateTimeAndPressure(valves map[string]valve) {
	s.timeRemaining--
	for _, valve := range s.open {
		s.pressure += valves[valve].flowRate
	}
}
