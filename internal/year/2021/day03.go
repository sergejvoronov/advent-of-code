package aoc2021

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day03 struct{}

func Day03() solution.Provider {
	return &day03{}
}

func (*day03) SolveA(input string) string {
	type report struct {
		gamma, epsilon string
	}

	numbers := strings.Split(input, "\n")
	r := report{}

	for x := 0; x < len(numbers[0]); x++ {
		zeros, ones := 0, 0
		for _, number := range numbers {
			if string(number[x]) == "0" {
				zeros++
				continue
			}
			ones++
		}

		if ones > zeros {
			r.gamma += "0"
			r.epsilon += "1"
			continue
		}
		r.gamma += "1"
		r.epsilon += "0"

	}

	gamma, _ := strconv.ParseInt(r.gamma, 2, 64)
	epsilon, _ := strconv.ParseInt(r.epsilon, 2, 64)

	return strconv.Itoa(int(gamma) * int(epsilon))
}

func (*day03) SolveB(input string) string {
	return solution.NotImplemented
}
