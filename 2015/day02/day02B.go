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
		var imax, maxEdge int
		bowLen := 1
		edges := make([]int, 3)

		size := strings.Split(s, "x")
		for k, v := range size {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			edges[k] = d
			bowLen *= d
		}

		for i, v := range edges {
			if maxEdge < v {
				maxEdge = v
				imax = i
			}
		}

		edges[imax] = edges[len(edges)-1]
		edges[len(edges)-1] = 0
		edges = edges[:len(edges)-1]

		answer += 2*edges[0] + 2*edges[1] + bowLen
	}

	return fmt.Sprintf("Day 02(B) answer: %d", answer)
}
