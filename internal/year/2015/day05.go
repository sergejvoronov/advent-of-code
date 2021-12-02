package aoc2015

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day05 struct{}

func Day05() solution.Provider {
	return &day05{}
}

func (*day05) SolveA(input string) string {
	var answer int
	xs := strings.Split(input, "\n")
	for _, s := range xs {
		if hasNoNaughtyStrings(s) && hasDoubleLetters(s) && hasThreeVowels(s) {
			answer++
		}
	}

	return strconv.Itoa(answer)
}

func hasNoNaughtyStrings(s string) bool {
	for _, v := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, v) {
			return false
		}
	}

	return true
}

func hasDoubleLetters(s string) bool {
	for k := range s {
		if k+1 < len(s) && s[k] == s[k+1] {
			return true
		}
	}

	return false
}

func hasThreeVowels(s string) bool {
	var contains int
	for _, v := range s {
		if strings.ContainsAny(string(v), "aeiou") {
			contains++
		}
		if contains == 3 {
			return true
		}
	}

	return false
}

func (*day05) SolveB(input string) string {
	var answer int
	xs := strings.Split(input, "\n")
	for _, s := range xs {
		if hasLetterBetween(s) && hasPairOfLetters(s) {
			answer++
		}
	}

	return strconv.Itoa(answer)
}

func hasLetterBetween(s string) bool {
	for k := range s {
		if k+2 < len(s) && s[k] == s[k+2] {
			return true
		}
	}

	return false
}

func hasPairOfLetters(s string) bool {
	size := len(s)
	for x := 0; x < size; x++ {
		for y := x + 2; y < size; y++ {
			if y+1 < size && s[x] == s[y] && s[x+1] == s[y+1] {
				return true
			}
		}
	}
	return false
}
