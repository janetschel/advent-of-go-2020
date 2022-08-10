package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"math"
	"strconv"
	"strings"
)

var boardSize int = 10
var rollsPerTurn int = 3

type player struct {
	position int
	score    int
}

type determininsticDie struct {
	values       []int
	currentIndex int
}

type game struct {
	p1 player
	p2 player
}

func main() {
	input := files.ReadFile(21, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	p1, p2 := parseInput(input)
	die := getDeterministicDie(1, 100)
	rolls := playGameDeterministically(&p1, &p2, &die)

	losingScore := p1.score
	if p2.score < p1.score {
		losingScore = p2.score
	}

	return rolls * losingScore
}

func solvePart2(input []string) int64 {
	p1, p2 := parseInput(input)
	r1, r2 := playGameQuantumly(p1, p2)
	return int64(math.Max(float64(r1), float64(r2)))
}

func playGameDeterministically(p1 *player, p2 *player, die *determininsticDie) int {
	turn := 0

	for p1.score < 1000 && p2.score < 1000 {
		rollTotal := slices.Sum(die.roll(rollsPerTurn))
		if turn%2 == 0 {
			p1.moveAndScore(rollTotal)
		} else {
			p2.moveAndScore(rollTotal)
		}
		turn++
	}

	return turn * rollsPerTurn
}

func playGameQuantumly(p1 player, p2 player) (int64, int64) {
	games := map[game]int64{{p1: p1, p2: p2}: 1}
	rollOptions := map[int]int64{}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rollOptions[i+j+k]++
			}
		}
	}

	var p1Wins, p2Wins int64 = 0, 0
	turn := 0

	for len(games) > 0 {
		next := map[game]int64{}
		for g, gameCount := range games {
			for roll, rollCount := range rollOptions {
				n := gameCount * rollCount
				if turn%2 == 0 {
					p1Copy := g.p1.moveAndScoreCopy(roll)
					if p1Copy.score >= 21 {
						p1Wins += n
					} else {
						next[game{p1: p1Copy, p2: g.p2}] += n
					}
				} else {
					p2Copy := g.p2.moveAndScoreCopy(roll)
					if p2Copy.score >= 21 {
						p2Wins += n
					} else {
						next[game{p1: g.p1, p2: p2Copy}] += n
					}
				}
			}
		}
		games = next
		turn++
	}

	return p1Wins, p2Wins
}

func parseInput(input []string) (player, player) {
	p1Position, _ := strconv.Atoi(strings.Split(input[0], ": ")[1])
	p2Position, _ := strconv.Atoi(strings.Split(input[1], ": ")[1])
	return player{position: p1Position, score: 0}, player{position: p2Position, score: 0}
}

func getDeterministicDie(min int, max int) determininsticDie {
	values := []int{}
	for i := min; i <= max; i++ {
		values = append(values, i)
	}
	return determininsticDie{values: values, currentIndex: 0}
}

func (p *player) moveAndScore(steps int) {
	p.position = ((p.position + steps - 1) % boardSize) + 1
	p.score += p.position
}

func (p *player) moveAndScoreCopy(steps int) player {
	position := ((p.position + steps - 1) % boardSize) + 1
	score := position + p.score
	return player{position: position, score: score}
}

func circularIncrement(index int, size int) int {
	return (index + 1) % size
}

func (d *determininsticDie) roll(rolls int) []int {
	r := []int{}
	for i := 0; i < rolls; i++ {
		r = append(r, d.values[d.currentIndex])
		d.currentIndex = circularIncrement(d.currentIndex, len(d.values))
	}
	return r
}
