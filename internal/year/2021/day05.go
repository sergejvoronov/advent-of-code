package aoc2021

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

const ventMapSize = 1000

type day05 struct{}

func Day05() solution.Provider {
	return &day05{}
}

func (*day05) SolveA(input string) string {
	var (
		ventMap [][]int
		answer  int
	)

	for x := 0; x < ventMapSize; x++ {
		ventMap = append(ventMap, make([]int, ventMapSize))
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		startCoords := strings.Split(coords[0], ",")
		endCoords := strings.Split(coords[1], ",")

		if startCoords[0] == endCoords[0] {
			x, _ := strconv.Atoi(startCoords[0])
			y1, _ := strconv.Atoi(startCoords[1])
			y2, _ := strconv.Atoi(endCoords[1])
			if y1 > y2 {
				for i := y2; i <= y1; i++ {
					ventMap[i][x]++
				}
				continue
			}

			for i := y1; i <= y2; i++ {
				ventMap[i][x]++
			}
			continue
		}

		if startCoords[1] == endCoords[1] {
			y, _ := strconv.Atoi(startCoords[1])
			x1, _ := strconv.Atoi(startCoords[0])
			x2, _ := strconv.Atoi(endCoords[0])
			if x1 > x2 {
				for i := x2; i <= x1; i++ {
					ventMap[y][i]++
				}
				continue
			}

			for i := x1; i <= x2; i++ {
				ventMap[y][i]++
			}
		}
	}

	for _, line := range ventMap {
		for _, point := range line {
			if point > 1 {
				answer++
			}
		}
	}

	return strconv.Itoa(answer)
}

func (*day05) SolveB(input string) string {
	var (
		ventMap [][]int
		answer  int
	)

	for x := 0; x < ventMapSize; x++ {
		ventMap = append(ventMap, make([]int, ventMapSize))
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		startCoords := strings.Split(coords[0], ",")
		endCoords := strings.Split(coords[1], ",")

		x1, _ := strconv.Atoi(startCoords[0])
		x2, _ := strconv.Atoi(endCoords[0])
		y1, _ := strconv.Atoi(startCoords[1])
		y2, _ := strconv.Atoi(endCoords[1])

		if startCoords[0] == endCoords[0] {
			x, _ := strconv.Atoi(startCoords[0])
			if y1 > y2 {
				for i := y2; i <= y1; i++ {
					ventMap[i][x]++
				}
				continue
			}

			for i := y1; i <= y2; i++ {
				ventMap[i][x]++
			}
			continue
		}

		if startCoords[1] == endCoords[1] {
			y, _ := strconv.Atoi(startCoords[1])
			if x1 > x2 {
				for i := x2; i <= x1; i++ {
					ventMap[y][i]++
				}
				continue
			}

			for i := x1; i <= x2; i++ {
				ventMap[y][i]++
			}
			continue
		}

		if x1 > x2 && y1 > y2 {
			for x, y := x2, y2; x <= x1; x, y = x+1, y+1 {
				ventMap[y][x]++
			}
			continue
		}

		if x1 < x2 && y1 < y2 {
			for x, y := x1, y1; x <= x2; x, y = x+1, y+1 {
				ventMap[y][x]++
			}
			continue
		}

		if x1 < x2 && y1 > y2 {
			for x, y := x1, y1; x <= x2; x, y = x+1, y-1 {
				ventMap[y][x]++
			}
			continue
		}

		if x1 > x2 && y1 < y2 {
			for x, y := x1, y1; y <= y2; x, y = x-1, y+1 {
				ventMap[y][x]++
			}
			continue
		}
	}

	for _, line := range ventMap {
		for _, point := range line {
			if point > 1 {
				answer++
			}
		}
	}

	return strconv.Itoa(answer)

}
