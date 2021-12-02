package aoc2020

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day02 struct{}

func Day02() solution.Provider {
	return &day02{}
}

func (*day02) SolveA(input string) string {
	var answer int

	type pack struct {
		minRequired, maxRequired int
		requiredLetter, password string
	}

	lines := strings.Split(input, "\n")
	packs := make([]pack, 0, len(lines))
	for _, v := range lines {
		line := strings.Split(v, ": ")
		policy := strings.Split(line[0], " ")
		minMax := strings.Split(policy[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		packs = append(packs, pack{
			minRequired:    min,
			maxRequired:    max,
			requiredLetter: policy[1],
			password:       line[1],
		})
	}

	for _, pack := range packs {
		letterAppeared := 0

		for _, letter := range pack.password {
			if string(letter) != pack.requiredLetter {
				continue
			}

			letterAppeared++
			if letterAppeared > pack.maxRequired {
				break
			}
		}
		if letterAppeared < pack.minRequired || letterAppeared > pack.maxRequired {
			continue
		}
		answer++
	}

	return strconv.Itoa(answer)
}

func (*day02) SolveB(input string) string {
	var answer int

	type pack struct {
		position1, position2 int
		letter, password     string
	}

	lines := strings.Split(input, "\n")
	packs := make([]pack, 0, len(lines))
	for _, v := range lines {
		line := strings.Split(v, ": ")
		policy := strings.Split(line[0], " ")
		positions := strings.Split(policy[0], "-")
		pos1, _ := strconv.Atoi(positions[0])
		pos2, _ := strconv.Atoi(positions[1])
		packs = append(packs, pack{
			position1: pos1 - 1,
			position2: pos2 - 1,
			letter:    policy[1],
			password:  line[1],
		})
	}

	for _, pack := range packs {
		if string(pack.password[pack.position1]) != pack.letter {
			if string(pack.password[pack.position2]) == pack.letter {
				answer++
			}
			continue
		}

		if string(pack.password[pack.position2]) != pack.letter {
			answer++
		}
	}

	return strconv.Itoa(answer)
}
