package aoc2021

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
	type position struct {
		horizontal, depth int
	}

	commands := strings.Split(input, "\n")

	p := position{}
	for _, command := range commands {
		split := strings.Split(command, " ")
		number, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			p.horizontal += number
		case "up":
			p.depth -= number
		case "down":
			p.depth += number
		}
	}

	return strconv.Itoa(p.horizontal * p.depth)
}

func (*day02) SolveB(input string) string {
	type position struct {
		horizontal, depth, aim int
	}

	commands := strings.Split(input, "\n")

	p := position{}
	for _, command := range commands {
		split := strings.Split(command, " ")
		number, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			p.horizontal += number
			p.depth += p.aim * number
		case "up":
			p.aim -= number
		case "down":
			p.aim += number
		}
	}

	return strconv.Itoa(p.horizontal * p.depth)
}
