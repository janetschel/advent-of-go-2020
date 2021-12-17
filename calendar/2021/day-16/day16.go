package main

import (
	"advent-of-go/utils/files"
	"fmt"
	"math"
	"strconv"
)

type packet struct {
	version      int64
	typeID       int64
	literal      int64
	lengthTypeID int
	length       int64
	subpackets   []packet
}

func main() {
	input := files.ReadFile(16, 2021, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int64 {
	binary := hexToBinary(input[0])
	_, _, sum := parsePacket(binary, 0, 0)

	return sum
}

func solvePart2(input []string) int64 {
	binary := hexToBinary(input[0])
	p, _, _ := parsePacket(binary, 0, 0)
	return executePacket(p)
}

func hexToBinary(hex string) string {
	binary := ""
	for i := 0; i < len(hex); i++ {
		asInt, _ := strconv.ParseUint(hex[i:i+1], 16, 64)
		binary += fmt.Sprintf("%04b", asInt)
	}
	return binary
}

func parseHeader(binary string, i int) (packet, int) {
	version, _ := strconv.ParseInt(binary[i:i+3], 2, 64)
	typeID, _ := strconv.ParseInt(binary[i+3:i+6], 2, 64)

	p := packet{
		version:      version,
		typeID:       typeID,
		literal:      -1,
		lengthTypeID: -1,
		length:       -1,
		subpackets:   nil,
	}
	return p, i + 6
}

func parseLiteral(binary string, i int) (int64, int) {
	lastPacket := false
	literalBinary := ""
	for !lastPacket {
		literalBinary += binary[i+1 : i+5]
		if binary[i:i+1] == "0" {
			lastPacket = true
		}
		i += 5
	}
	l, _ := strconv.ParseInt(literalBinary, 2, 64)
	return l, i
}

func executePacket(p packet) int64 {
	switch p.typeID {
	case int64(0):
		sum := int64(0)
		for i := range p.subpackets {
			sum += executePacket(p.subpackets[i])
		}
		return sum
	case int64(1):
		product := int64(1)
		for i := range p.subpackets {
			product *= executePacket(p.subpackets[i])
		}
		return product
	case int64(2):
		min := int64(math.MaxInt)
		for i := range p.subpackets {
			result := executePacket(p.subpackets[i])
			if result < min {
				min = result
			}
		}
		return min
	case int64(3):
		max := int64(0)
		for i := range p.subpackets {
			result := executePacket(p.subpackets[i])
			if result > max {
				max = result
			}
		}
		return max
	case int64(4):
		return p.literal
	case int64(5):
		if executePacket(p.subpackets[0]) > executePacket(p.subpackets[1]) {
			return 1
		}
		return 0
	case int64(6):
		if executePacket(p.subpackets[0]) < executePacket(p.subpackets[1]) {
			return 1
		}
		return 0
	case int64(7):
		if executePacket(p.subpackets[0]) == executePacket(p.subpackets[1]) {
			return 1
		}
		return 0
	}
	return p.literal
}

func parsePacket(binary string, i int, runningSum int64) (packet, int, int64) {
	p, i := parseHeader(binary, i)

	runningSum += p.version

	if p.typeID == 4 {
		p.literal, i = parseLiteral(binary, i)
	} else {
		i++
		if binary[i-1:i] == "1" {
			p.lengthTypeID = 1
			subPackets, _ := strconv.ParseInt(binary[i:i+11], 2, 64)
			p.length = subPackets
			p.subpackets = []packet{}
			i += 11
			for s := 0; s < int(subPackets) && len(binary[i:]) > 6; s++ {
				sub := packet{}
				sub, i, runningSum = parsePacket(binary, i, runningSum)
				p.subpackets = append(p.subpackets, sub)
			}
		} else {
			p.lengthTypeID = 0
			length, _ := strconv.ParseInt(binary[i:i+15], 2, 64)
			p.length = length
			p.subpackets = []packet{}
			i += 15
			start := i
			for i-start < int(length) && len(binary[i:]) > 6 {
				sub := packet{}
				sub, i, runningSum = parsePacket(binary, i, runningSum)
				p.subpackets = append(p.subpackets, sub)
			}
		}
	}
	return p, i, runningSum
}
