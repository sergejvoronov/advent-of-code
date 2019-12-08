package day01

import (
	"fmt"
)

func SolveA(input string) string {
	var answer int

	for _, f := range input {
		switch f {
		case '(':
			answer++
		case ')':
			answer--
		}
	}

	return fmt.Sprintf("Day 01(A) answer: %d", answer)
}
