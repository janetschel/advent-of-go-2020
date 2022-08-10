package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"advent-of-go/utils/sets"
	"fmt"
	"strconv"
	"strings"
)

type scanner struct {
	label       string
	beacons     []coords
	position    coords
	distanceMap map[int][]*coords
}

type coords struct {
	x, y, z int
}

func main() {
	input := files.ReadFile(19, 2021, "\n\n")
	scanners := parseInput(input)
	normalized := normalizeScanners(scanners, 12)
	println(solvePart1(normalized))
	println(solvePart2(normalized))
}

func solvePart1(normalized []scanner) int {
	keySet := sets.New()
	for s := range normalized {
		for b := range normalized[s].beacons {
			keySet.Add(beaconKey(normalized[s].beacons[b]))
		}
	}

	return keySet.Size()
}

func solvePart2(normalized []scanner) int {
	maxDistance := 0
	for i := range normalized {
		for j := i + 1; j < len(normalized); j++ {
			dist := distance(normalized[i].position, normalized[j].position)
			if dist > maxDistance {
				maxDistance = dist
			}
		}
	}
	return maxDistance
}

func normalizeScanners(scanners []scanner, minScanners int) []scanner {
	normalized, remaining := scanners[:1], scanners[1:]
	lastNormalizedIndex := 0
	killSwitch := 0
	for len(normalized) < len(scanners) && killSwitch < 100000 {
		killSwitch++
		if lastNormalizedIndex >= len(normalized) {
			printNormalizedBreakdown(normalized, remaining)
			panic(fmt.Sprintf("Failed to normalize any additional scanners"))
		}
		anchor := normalized[lastNormalizedIndex]
		for i := 0; i < len(remaining); i++ {
			s, couldNormalize := normalizeScanner(&anchor, &remaining[i], minScanners)
			if couldNormalize {
				normalized = append(normalized, *s)
				temp := make([]scanner, len(remaining))
				copy(temp, remaining)
				remaining = append(temp[:i], temp[i+1:]...)
				fmt.Printf("Normalized beacon %v to beacon %v, %v beacons remaining\n", s.label, anchor.label, len(remaining))
				printNormalizedBreakdown(normalized, remaining)
				fmt.Println()
				i--
			}
		}
		lastNormalizedIndex++
	}
	return normalized
}

func normalizeScanner(normalized *scanner, toNormalize *scanner, minScanners int) (*scanner, bool) {
	if !scannersOverlap(*normalized, *toNormalize, minScanners) {
		return toNormalize, false
	}

	rotations, translations := map[int]int{}, map[string]int{}
	for dist, normalizedPair := range normalized.distanceMap {
		pairToTransform, inMap := toNormalize.distanceMap[dist]
		if inMap {
			beacon1Rotations, beacon2Rotations := getRotations(*pairToTransform[0]), getRotations(*pairToTransform[1])
			for i, r1 := range beacon1Rotations {
				for j, r2 := range beacon2Rotations {
					v1, v2 := subtract(*normalizedPair[0], r1), subtract(*normalizedPair[1], r2)
					if equals(v1, v2) {
						rotations[i]++
						rotations[j]++
						translations[beaconKey(v1)]++
					}
				}
			}
		}
	}
	bestRotation := 0
	for r, count := range rotations {
		if count > rotations[bestRotation] && count >= 12 {
			bestRotation = r
		}
	}
	bestTranslation := ""
	for t, count := range translations {
		if count > translations[bestTranslation] && count >= 12 {
			bestTranslation = t
		}
	}
	if bestRotation == 0 || bestTranslation == "" {
		return toNormalize, false
	}

	translation := coordsFromKey(bestTranslation)
	normalizedBeacons := []coords{}
	for i := range toNormalize.beacons {
		normalizedBeacons = append(normalizedBeacons, transform(toNormalize.beacons[i], translation, bestRotation))
	}
	toNormalize.beacons = normalizedBeacons
	s := scanner{
		label:    toNormalize.label,
		beacons:  normalizedBeacons,
		position: transform(coords{x: 0, y: 0}, translation, bestRotation),
	}
	initializeDistanceMaps(&s)
	return &s, true
}

func parseInput(input []string) []scanner {
	scanners := []scanner{}
	for i := range input {
		scanner := scanner{label: fmt.Sprintf("%v", i)}
		beacons := strings.Split(input[i], "\n")[1:]
		for _, line := range beacons {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			z, _ := strconv.Atoi(parts[2])
			beacon := coords{x: x, y: y, z: z}
			scanner.beacons = append(scanner.beacons, beacon)
		}
		initializeDistanceMaps(&scanner)
		scanners = append(scanners, scanner)
	}
	return scanners
}

func distance(c1 coords, c2 coords) int {
	return maths.Abs(c1.x-c2.x) + maths.Abs(c1.y-c2.y) + maths.Abs(c1.z-c2.z)
}

func add(c1 coords, c2 coords) coords {
	return coords{x: c1.x + c2.x, y: c1.y + c2.y, z: c1.z + c2.z}
}

func subtract(c1 coords, c2 coords) coords {
	return coords{x: c1.x - c2.x, y: c1.y - c2.y, z: c1.z - c2.z}
}

func equals(c1 coords, c2 coords) bool {
	return c1.x == c2.x && c1.y == c2.y && c1.z == c2.z
}

func transform(c coords, translation coords, rotation int) coords {
	return add(getRotations(c)[rotation], translation)
}

func scannersOverlap(s1 scanner, s2 scanner, minBeacons int) bool {
	overlap := 0
	for dist := range s1.distanceMap {
		_, inMap := s2.distanceMap[dist]
		if inMap {
			overlap++
		}
	}
	return overlap >= minBeacons
}

func getRotations(c coords) []coords {
	return []coords{
		{x: c.x, y: c.y, z: c.z},
		{c.x, -c.z, c.y},
		{c.x, -c.y, -c.z},
		{c.x, c.z, -c.y},

		{-c.y, c.x, c.z},
		{c.z, c.x, c.y},
		{c.y, c.x, -c.z},
		{-c.z, c.x, -c.y},

		{-c.x, -c.y, c.z},
		{-c.x, -c.z, -c.y},
		{-c.x, c.y, -c.z},
		{-c.x, c.z, c.y},

		{c.y, -c.x, c.z},
		{c.z, -c.x, -c.y},
		{-c.y, -c.x, -c.z},
		{-c.z, -c.x, c.y},

		{-c.z, c.y, c.x},
		{c.y, c.z, c.x},
		{c.z, -c.y, c.x},
		{-c.y, -c.z, c.x},

		{-c.z, -c.y, -c.x},
		{-c.y, c.z, -c.x},
		{c.z, c.y, -c.x},
		{c.y, -c.z, -c.x},
	}
}

func initializeDistanceMaps(s *scanner) {
	s.distanceMap = map[int][]*coords{}
	for i := range s.beacons {
		for j := i + 1; j < len(s.beacons); j++ {
			d := distance(s.beacons[i], s.beacons[j])
			s.distanceMap[d] = []*coords{&s.beacons[i], &s.beacons[j]}
		}
	}
}

func beaconKey(c coords) string {
	return fmt.Sprintf("%v,%v,%v", c.x, c.y, c.z)
}

func coordsFromKey(key string) coords {
	parts := strings.Split(key, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return coords{x: x, y: y, z: z}
}

func printNormalizedBreakdown(normalized []scanner, remaining []scanner) {
	for i := range normalized {
		fmt.Printf("%v ", normalized[i].label)
	}
	fmt.Print("| ")
	for i := range remaining {
		fmt.Printf("%v ", remaining[i].label)
	}
	fmt.Println()
}
