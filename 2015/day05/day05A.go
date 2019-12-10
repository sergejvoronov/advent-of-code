package day05

import (
	"fmt"
	"strings"
)

func SolveA(input string) string {
	var answer int
	xs := strings.Split(input, "\n")
	for _, s := range xs {
		if s == "" {
			continue
		}
		if hasNoNaughtyStrings(s) && hasDoubleLetters(s) && hasThreeVowels(s) {
			answer++
		}
	}

	return fmt.Sprintf("Day 05(A) answer: %d", answer)
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
