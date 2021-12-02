package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day04 struct{}

func Day04() solution.Provider {
	return &day04{}
}

func (*day04) SolveA(input string) string {
	var answer int

	entity := ""
	entities := make([]string, 0)

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		if s.Text() == "" {
			entities = append(entities, entity)
			entity = ""
		}
		entity += s.Text() + " "
	}

	for _, entity := range entities {
		passport := make(map[string]string)
		fields := strings.Fields(entity)

		for _, field := range fields {
			split := strings.Split(field, ":")
			passport[split[0]] = split[1]
		}
		if len(passport) == 8 {
			answer++
		}
		if len(passport) == 7 {
			if _, ok := passport["cid"]; !ok {
				answer++
			}
		}
	}

	return strconv.Itoa(answer)
}

func (*day04) SolveB(input string) string {
	var answer int

	// 	input = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
	// byr:1937 iyr:2017 cid:147 hgt:183cm

	// iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
	// hcl:#cfa07d byr:1929

	// hcl:#ae17e1 iyr:2013
	// eyr:2024
	// ecl:brn pid:760753108 byr:1931
	// hgt:179cm

	// hcl:#cfa07d eyr:2025 pid:166559648
	// iyr:2011 ecl:brn hgt:59in
	// `
	entity := ""
	entities := make([]string, 0)
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		if s.Text() == "" {
			entities = append(entities, entity)
			entity = ""
		}
		entity += s.Text() + " "
	}

	for _, entity := range entities {
		passport := make(map[string]string)
		fields := strings.Fields(entity)

		for _, field := range fields {
			split := strings.Split(field, ":")
			passport[split[0]] = split[1]
		}
		if len(passport) == 8 {
			answer++
		}
		if len(passport) == 7 {
			if _, ok := passport["cid"]; !ok {
				answer++
			}
		}
	}

	return solution.NotImplemented
}
