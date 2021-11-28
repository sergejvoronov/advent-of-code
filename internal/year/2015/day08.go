package aoc2015

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day08 struct{}

func Day08() solution.Provider {
	return &day08{}
}

func (*day08) SolveA(input string) string {
	charCount, inMemory := 0, 0
	xs := strings.Split(input, "\n")
	for _, s := range xs {
		if s == "" {
			continue
		}
		charCount += len(s)
		inMemory += utf8.RuneCountInString(s)
	}
	return strconv.Itoa(charCount - inMemory)
}

func (*day08) SolveB(input string) string {
	return solution.NotImplemented
}
