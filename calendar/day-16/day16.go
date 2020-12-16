package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	"advent-of-go-2020/utils/sets"
	"strings"
)

type Range struct {
	from int
	to int
}

func main() {
	input := files.ReadFile(16, "\n\n")
	result, resultPart2 := solve(input)

	println("Solution part 1:", result)
	println("Solution part 2:", resultPart2)
}

func solve(input []string) (int, int) {
	ownTicket := strings.Split(input[1],"\n")[1]
	nearbyTickets := strings.Split(input[2], "\n")[1:]
	rules := parseRules(input)

	result, validTickets := solvePart1(rules, nearbyTickets)
	resultPart2 := solvePart2(rules, validTickets, ownTicket)

	return result, resultPart2
}

func solvePart1(all map[string][]Range, nearbyTickets []string) (int, []string) {
	result, validTickets, valid := 0, make([]string, 0), false

	for _, ticket := range nearbyTickets {
		for _, field := range conv.ToIntSlice(strings.Split(ticket, ",")) {
			valid = false

		out:
			for _, rules := range all {
				for _, rule := range rules {
					if valid = isValid(rule, field); valid {
						break out
					}
				}
			}

			if !valid {
				result += field
				break
			}
		}

		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	return result, validTickets
}

func solvePart2(all map[string][]Range, valid []string, myTicket string) int {
	myTicketFields := strings.Split(myTicket, ",")
	result := 1
	seenValids := sets.New()

	for i := 0; i < len(myTicketFields); i++ {
		possibleValids := make(map[string]int)

		for j := 0; j < len(valid); j++ {
			fields := conv.ToIntSlice(strings.Split(valid[j], ","))

			for name, rules := range all {
				for _, rule := range rules {
					if isValid(rule, fields[i]) {
						possibleValids[name] += 1
					}
				}
			}
		}

		for name, value := range possibleValids {
			if !seenValids.Has(name) && value == len(valid) {
				result *= conv.ToInt(myTicketFields[i])
				seenValids.Add(name)
			}
		}

	}

	return result
}

func isValid(rule Range, field int) bool {
	return field >= rule.from && field <= rule.to
}

func parseRules(input []string) map[string][]Range {
	rules, valid := strings.Split(input[0], "\n"), make(map[string][]Range)

	for _, element := range rules {
		split := strings.Split(element, ": ")

		for _, currRange := range strings.Split(split[1], " or ") {
			nums := conv.ToIntSlice(strings.Split(currRange, "-"))
			valid[split[0]] = append(valid[split[0]], Range {
				from: nums[0],
				to:   nums[1],
			})
		}
	}

	return valid
}
