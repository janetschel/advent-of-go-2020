package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"fmt"
	"strconv"
	"strings"
)

type valve struct {
	name string
	flowRate int
	tunnelsTo []string
	paths map[string]int
}

type state struct {
	current string
	timeRemaining int
	open map[string]int
}

func main() {
	input := files.ReadFile(16, 2022, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	valves := parseInput(input)
	pressures := getMaxPressureReleased(valves, 30, map[string]int{})
	maxPressure := 0
	for p := range pressures {
		if p > maxPressure {
			maxPressure = p
		}
	}
	return maxPressure
}

func solvePart2(input []string) int {
	valves := parseInput(input)
	pressures := getMaxPressureReleased(valves, 26, map[string]int{})
	maxPressure := 0
	for p := range pressures {
		if p > maxPressure {
			maxPressure = p
		}
	}
	return maxPressure
}

func getMaxPressureReleased(valves map[string]*valve, timeLimit int, open map[string]int) map[int]map[string]int {
	starting := state{ current: "AA", timeRemaining: timeLimit, open: open }
	queue := []*state{ &starting }

	endingPressures := map[int]map[string]int{}

	killswitch := 0
	for len(queue) > 0 && killswitch < 9999999 {
		currentState := queue[0]
		queue = queue[1:]

		if currentState.timeRemaining <= 0 {
			endingPressures[currentState.getEndPressure(valves)] = currentState.open
		} else {
			generatedStates := 0
			for v, pathLength := range valves[currentState.current].paths {
				_, open := currentState.open[v]
				// visit any/all unopened valves
				if !open && valves[v].flowRate > 0 {
					newState := currentState.copy()
					newState.current = v
					if newState.timeRemaining - (pathLength + 1) >= 0 {
						newState.timeRemaining -= pathLength + 1
						newState.open[v] = newState.timeRemaining
						queue = append(queue, &newState)
					} else {
						endingPressures[currentState.getEndPressure(valves)] = currentState.open
					}
					generatedStates++
				}
			}

			// if no moves/opens are relevant, stay in place
			if generatedStates == 0 {
				endingPressures[currentState.getEndPressure(valves)] = currentState.open
			}
		}
		killswitch++
	}

	println(killswitch, len(queue), len(endingPressures))
	return endingPressures
}

func parseInput(input []string) (map[string]*valve) {
	valves := map[string]*valve{}
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
		valves[parts[1]] = &valve
	}
	
	for _, v := range valves {
		v.generateShortestPaths(valves)
	}
	return valves
}

func (v *valve) generateShortestPaths(valves map[string]*valve) {
	queue := []string{v.name}
	v.paths = map[string]int{}
	visited := sets.New()
	visited.Add(v.name)

	for len(queue) > 0 {
		current := valves[queue[0]]
		queue = queue[1:]
		for _, next := range current.tunnelsTo {
			if !visited.Has(next) {
				visited.Add(next)
				v.paths[next] += v.paths[current.name] + 1
				queue = append(queue, next)
			}
		}
	}
}



func (s *state) getEndPressure(valves map[string]*valve) int {
	pressure := 0
	for v, t := range s.open {
		pressure += t * valves[v].flowRate
	}
	return pressure
}

func (s state) ToString() string {
	return fmt.Sprintf("%s %d %v", s.current, s.timeRemaining, s.open)
}

func (s state) copy() state {
	newOpen := map[string]int{}
	for v, t := range s.open {
		newOpen[v] = t
	}
	return state{
		current: s.current,
		timeRemaining: s.timeRemaining,
		// pressure: s.pressure,
		open: newOpen,
	}
}
