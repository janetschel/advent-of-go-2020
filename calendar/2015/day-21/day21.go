package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/slices"
	"math"
	"strconv"
	"strings"
)

type item struct {
	cost   int
	damage int
	armor  int
}

type stats struct {
	hitPoints int
	damage    int
	armor     int
}

type shop struct {
	weapons []item
	armor   []item
	rings   []item
}

func main() {
	input := files.ReadFile(21, 2015, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	minGold := math.MaxInt
	blocks := buildStatBlocks()
	boss := parseBoss(input)
	for block, gold := range blocks {
		if gold < minGold && simulateGame(*block, boss) {
			minGold = gold
		}
	}

	return minGold
}

func solvePart2(input []string) int {
	maxGold := 0
	blocks := buildStatBlocks()
	boss := parseBoss(input)
	for block, gold := range blocks {
		if gold > maxGold && !simulateGame(*block, boss) {
			maxGold = gold
		}
	}

	return maxGold
}

func parseBoss(input []string) stats {
	hp, _ := strconv.Atoi(strings.Split(input[0], ": ")[1])
	damage, _ := strconv.Atoi(strings.Split(input[1], ": ")[1])
	armor, _ := strconv.Atoi(strings.Split(input[2], ": ")[1])
	return stats{
		hitPoints: hp,
		damage:    damage,
		armor:     armor,
	}
}

func calculateDamage(damageScore int, armorScore int) int {
	baseDamage := damageScore - armorScore
	if baseDamage > 0 {
		return baseDamage
	}
	return 1
}

func buildShop() shop {
	weapons := []item{
		{cost: 8, damage: 4, armor: 0},
		{cost: 10, damage: 5, armor: 0},
		{cost: 25, damage: 6, armor: 0},
		{cost: 40, damage: 7, armor: 0},
		{cost: 74, damage: 8, armor: 0},
	}
	armor := []item{
		{cost: 13, damage: 0, armor: 1},
		{cost: 31, damage: 0, armor: 2},
		{cost: 53, damage: 0, armor: 3},
		{cost: 75, damage: 0, armor: 4},
		{cost: 102, damage: 0, armor: 5},
	}
	rings := []item{
		{cost: 25, damage: 1, armor: 0},
		{cost: 50, damage: 2, armor: 0},
		{cost: 100, damage: 3, armor: 0},
		{cost: 20, damage: 0, armor: 1},
		{cost: 40, damage: 0, armor: 2},
		{cost: 80, damage: 0, armor: 3},
	}
	return shop{
		weapons: weapons,
		armor:   armor,
		rings:   rings,
	}
}

func simulateGame(playerParam stats, bossParam stats) bool {
	player, boss := playerParam, bossParam
	for i := 0; player.hitPoints > 0 && boss.hitPoints > 0; i++ {
		if i%2 == 0 {
			damage := calculateDamage(player.damage, boss.armor)
			boss.hitPoints -= damage
		} else {
			damage := calculateDamage(boss.damage, player.armor)
			player.hitPoints -= damage
		}
	}
	return player.hitPoints > 0
}

func buildStatBlocks() map[*stats]int {
	shop := buildShop()
	ringOptions := [][]item{{}}
	referenceSlice := []int{}
	for i := range shop.rings {
		referenceSlice = append(referenceSlice, i)
		ringOptions = append(ringOptions, []item{shop.rings[i]})
	}
	pairs := slices.GenerateCombinationsLengthN(referenceSlice, 2)
	for i := range pairs {
		ringOptions = append(ringOptions, []item{shop.rings[pairs[i][0]], shop.rings[pairs[i][1]]})
	}

	blocks := map[*stats]int{}
	for _, weapon := range shop.weapons {
		for armor := -1; armor < len(shop.armor); armor++ {
			for _, rings := range ringOptions {
				gold := weapon.cost
				damageMod := weapon.damage

				armorMod := 0
				if armor >= 0 {
					armorMod += shop.armor[armor].armor
					gold += shop.armor[armor].cost
				}
				for _, ring := range rings {
					armorMod += ring.armor
					damageMod += ring.damage
					gold += ring.cost
				}
				block := stats{
					hitPoints: 100,
					armor:     armorMod,
					damage:    damageMod,
				}
				blocks[&block] = gold
			}
		}
	}

	return blocks
}
