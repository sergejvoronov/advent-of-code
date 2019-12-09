package day03

import (
	"fmt"
)

func SolveA(input string) string {
	type House struct {
		X, Y int
	}

	h := House{0, 0}
	presents := make(map[House]int)
	presents[h]++
	for _, direction := range input {
		switch string(direction) {
		case ">":
			h.X++
		case "<":
			h.X--
		case "^":
			h.Y++
		case "v":
			h.Y--
		default:
			continue
		}
		presents[h]++
	}

	return fmt.Sprintf("Day 03(A) answer: %d", len(presents))
}
