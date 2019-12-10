package day05

import (
	"fmt"
	"strings"
)

func SolveB(input string) string {
	var answer int
	xs := strings.Split(input, "\n")
	for _, s := range xs {
		if s == "" {
			continue
		}
		if hasLetterBetween(s) && hasPairOfLetters(s) {
			answer++
		}
	}

	return fmt.Sprintf("Day 05(B) answer: %d", answer)
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
