package aoc2015

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day02 struct{}

func Day02() solution.Provider {
	return &day02{}
}

func (*day02) SolveA(input string) string {
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

	return strconv.Itoa(answer)
}

func (*day02) SolveB(input string) string {
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

	return strconv.Itoa(answer)
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
