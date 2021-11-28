package aoc2015

import (
	"strconv"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day01 struct{}

func Day01() solution.Provider {
	return &day01{}
}

func (*day01) SolveA(input string) string {
	var answer int
	for _, f := range input {
		switch f {
		case '(':
			answer++
		case ')':
			answer--
		}
	}

	return strconv.Itoa(answer)
}

func (*day01) SolveB(input string) string {
	var answer int
	for k, f := range input {
		if answer == -1 {
			answer = k
			break
		}
		switch f {
		case '(':
			answer++
		case ')':
			answer--
		}
	}

	return strconv.Itoa(answer)
}
