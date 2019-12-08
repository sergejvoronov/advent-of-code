package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveB(input string) string {
	var answer int
	sizes := strings.Split(input, "\n")
	for _, s := range sizes {
		if s == "" {
			continue
		}

		bowLen := 1
		sides := make([]int, 3)
		sideLen := strings.Split(s, "x")
		for k, v := range sideLen {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			sides[k] = d
			bowLen *= d
		}

		sides = removeMaxSide(sides)

		answer += 2*sides[0] + 2*sides[1] + bowLen
	}

	return fmt.Sprintf("Day 02(B) answer: %d", answer)
}

func removeMaxSide(edges []int) []int {
	var maxSideIndex, maxSideValue int
	for i, v := range edges {
		if maxSideValue < v {
			maxSideValue = v
			maxSideIndex = i
		}
	}

	edges[maxSideIndex] = edges[len(edges)-1]
	edges[len(edges)-1] = 0
	edges = edges[:len(edges)-1]
	return edges
}
