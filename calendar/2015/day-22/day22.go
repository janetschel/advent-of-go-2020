package main

import (
	"advent-of-go/utils/files"
	"math"
	"strconv"
	"strings"
)

type gameState struct {
	bossHp     int
	bossDamage int

	playerHp int
	mana     int

	manaSpent int

	shieldDuration   int
	poisonDuration   int
	rechargeDuration int
}

func main() {
	input := files.ReadFile(22, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	initial := parseInitialGameState(input)
	game := findBestGameState(initial, false)
	return game.manaSpent
}

func solvePart2(input []string) int {
	initial := parseInitialGameState(input)
	game := findBestGameState(initial, true)
	return game.manaSpent
}

func gameOver(game gameState) bool {
	return game.bossHp <= 0 ||
		game.playerHp <= 0 ||
		game.mana <= 0
}

func playerWins(game gameState) bool {
	return gameOver(game) && game.bossHp <= 0
}

func spend(game *gameState, cost int) {
	game.mana -= cost
	game.manaSpent += cost
}

func magicMissile(game gameState) gameState {
	gs := game
	spend(&gs, 53)
	gs.bossHp -= 4
	return gs
}

func drain(game gameState) gameState {
	gs := game
	spend(&gs, 73)
	gs.bossHp -= 2
	gs.playerHp += 2
	return gs
}

func shield(game gameState) gameState {
	gs := game
	spend(&gs, 113)
	gs.shieldDuration = 6
	return gs
}

func poison(game gameState) gameState {
	gs := game
	spend(&gs, 173)
	gs.poisonDuration = 6
	return gs
}

func recharge(game gameState) gameState {
	gs := game
	spend(&gs, 229)
	gs.rechargeDuration = 5
	return gs
}

func playerArmor(game gameState) int {
	if game.shieldDuration > 0 {
		return 7
	}
	return 0
}

func resolveEffects(game gameState) gameState {
	gs := game

	if gs.shieldDuration > 0 {
		gs.shieldDuration--
	}

	if gs.poisonDuration > 0 {
		gs.bossHp -= 3
		gs.poisonDuration--
	}

	if gs.rechargeDuration > 0 {
		gs.mana += 101
		gs.rechargeDuration--
	}

	return gs
}

func takeBossTurn(game gameState) gameState {
	gs := game

	damage := gs.bossDamage - playerArmor(gs)
	if damage > 1 {
		gs.playerHp -= damage
	} else {
		gs.playerHp--
	}
	return gs
}

func playRound(game gameState, s spell, isPart2 bool) gameState {
	gs := game

	if isPart2 {
		gs.playerHp--
	}

	if gameOver(gs) {
		return gs
	}

	gs = resolveEffects(gs)

	if gameOver(gs) {
		return gs
	}

	gs = s(gs)

	if gameOver(gs) {
		return gs
	}

	// boss's turn starts with effects
	gs = resolveEffects(gs)

	if gameOver(gs) {
		return gs
	}

	gs = takeBossTurn(gs)

	if gameOver(gs) {
		return gs
	}

	return gs
}

func findBestGameState(initial gameState, isPart2 bool) gameState {
	nextUp := []gameState{initial}

	leastMana, bestGame := math.MaxInt, initial
	for len(nextUp) > 0 {
		current := nextUp[len(nextUp)-1]
		nextUp = nextUp[:len(nextUp)-1]

		if gameOver(current) {
			if playerWins(current) && current.manaSpent < leastMana {
				leastMana = current.manaSpent
				bestGame = current
			}
		} else {
			spells := getSpellOptions(current)
			for _, s := range spells {
				next := playRound(current, s, isPart2)
				if gameOver(next) {
					if playerWins(next) && next.manaSpent < leastMana {
						leastMana = next.manaSpent
						bestGame = next
					}
				} else {
					nextUp = append(nextUp, next)
				}
			}
		}
	}
	return bestGame
}

type spell func(gameState) gameState

func getSpellOptions(game gameState) []spell {
	manaAfterRecharge := game.mana
	if game.rechargeDuration > 0 {
		manaAfterRecharge += 101
	}
	spells := []spell{}
	if manaAfterRecharge >= 53 {
		spells = append(spells, magicMissile)
	}
	if manaAfterRecharge >= 73 {
		spells = append(spells, drain)
	}
	if game.shieldDuration <= 1 && manaAfterRecharge >= 113 {
		spells = append(spells, shield)
	}
	if game.poisonDuration <= 1 && manaAfterRecharge >= 173 {
		spells = append(spells, poison)
	}
	if game.rechargeDuration <= 1 && manaAfterRecharge >= 229 {
		spells = append(spells, recharge)
	}
	return spells
}

func parseInitialGameState(input []string) gameState {
	bossHp, _ := strconv.Atoi(strings.Split(input[0], ": ")[1])
	bossDamage, _ := strconv.Atoi(strings.Split(input[1], ": ")[1])
	return gameState{
		playerHp:   50,
		mana:       500,
		bossHp:     bossHp,
		bossDamage: bossDamage,
		manaSpent:  0,
	}
}
