package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveA(input string) string {
	var answer int
	edges := make([]int, 3)
	sizes := strings.Split(input, "\n")
	for _, s := range sizes {
		if s == "" {
			continue
		}

		var areas []int

		size := strings.Split(s, "x")
		for k, v := range size {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			edges[k] = d
		}

		areas = append(areas, edges[0]*edges[1])
		areas = append(areas, edges[1]*edges[2])
		areas = append(areas, edges[2]*edges[0])

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
