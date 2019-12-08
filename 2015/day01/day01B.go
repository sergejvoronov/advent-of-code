package day01

import (
	"fmt"
)

func SolveB(input string) string {
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

	return fmt.Sprintf("Day 01(B) answer: %d", answer)
}
