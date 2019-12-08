package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveA(input string) string {
	var answer int
	sides := make([]int, 3)
	sizes := strings.Split(input, "\n")
	for _, s := range sizes {
		if s == "" {
			continue
		}

		var areas []int

		sideLen := strings.Split(s, "x")
		for k, v := range sideLen {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			sides[k] = d
		}

		areas = append(areas, sides[0]*sides[1])
		areas = append(areas, sides[1]*sides[2])
		areas = append(areas, sides[2]*sides[0])

		extraPaper := areas[0]
		for _, v := range areas {
			if v < extraPaper {
				extraPaper = v
			}
		}

		answer += 2*(areas[0]+areas[1]+areas[2]) + extraPaper
	}

	return fmt.Sprintf("Day 02(A) answer: %d", answer)
}
