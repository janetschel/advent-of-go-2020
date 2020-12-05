package main

import (
	"advent-of-go-2020/utils/conv"
	"advent-of-go-2020/utils/files"
	maps "advent-of-go-2020/utils/map"
	"advent-of-go-2020/utils/regex"
	"strings"
)

type Passport struct {
	byr, iyr, eyr int
	hcl, ecl, pid string
	hgt Height
}

type Height struct {
	hgt int
	unit string
}

func (p *Passport) init(parts map[string]string) {
	p.byr = conv.ToIntOrElse(parts["byr"], 0)
	p.iyr = conv.ToIntOrElse(parts["iyr"], 0)
	p.eyr = conv.ToIntOrElse(parts["eyr"], 0)
	p.hcl = parts["hcl"]
	p.ecl = parts["ecl"]
	p.pid = parts["pid"]
	p.hgt = parseHeight(parts["hgt"])
}

func main() {
	input := files.ReadFile(4, "\n\n")
	println("Valid passports:", solveReworked(input))
}

func makeMap(element string) map[string]string {
	parts := make(map[string]string)

	for _, curr := range strings.Fields(element) {
		part := strings.Split(curr, ":")
		parts[part[0]] = part[1]
	}

	return parts
}

func parseHeight(hgt string) Height {
	if len(hgt) < 2 {
		return Height{ hgt:  0, unit: ""}
	}

	return Height {
		hgt:  conv.ToIntOrElse(hgt[:len(hgt) - 2], 0),
		unit: hgt[len(hgt) - 2:],
	}
}

func solveReworked(input []string) int {
	valid := 0

	for _, element := range input {
		passport := new(Passport)
		passport.init(makeMap(element))

		if validPassport(*passport) {
			valid++
		}
	}

	return valid
}

func validPassport(p Passport) bool {
	checks := map[string]bool {
		"byr": 1920 <= p.byr && p.byr <= 2002,
		"iyr": 2010 <= p.iyr && p.iyr <= 2020,
		"eyr": 2020 <= p.eyr && p.eyr <= 2030,
		"hgt": validHeight(p.hgt),
		"ecl": regex.Match(p.ecl, "\\b(amb|blu|brn|gry|grn|hzl|oth)\\b"),
		"hcl": regex.Match(p.hcl, "#[0-9a-f]{6}"),
		"pid": regex.Match(p.pid, "^\\d{9}$"),
	}

	valid := maps.All(checks, true)
	return valid
}

func validHeight(hgt Height) bool {
	if hgt.unit == "cm" && 150 <= hgt.hgt && hgt.hgt <= 193 {
		return true
	} else if hgt.unit == "in" && 59 <= hgt.hgt && hgt.hgt <= 76 {
		return true
	}

	return false
}
