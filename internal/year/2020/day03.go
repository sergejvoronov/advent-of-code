package aoc2020

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day03 struct{}

func Day03() solution.Provider {
	return &day03{}
}

func (*day03) SolveA(input string) string {
	var (
		answer      int
		right, down = 3, 1
		x, y        = 0, 0
	)

	grid := strings.Split(input, "\n")
	grid = grid[:len(grid)-1]

	xAxisLength := len(grid[0])
	yAxisLength := len(grid)

	for range grid {
		y += down
		x += right
		if x > xAxisLength-1 {
			x -= xAxisLength
		}
		if string(grid[y][x]) == "#" {
			answer++
		}
		if yAxisLength-1 == y {
			break
		}
	}

	return strconv.Itoa(answer)
}

func (*day03) SolveB(input string) string {
	var answer int = 1

	type slope struct {
		right, down int
	}

	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	grid := strings.Split(input, "\n")

	xAxisLength := len(grid[0])
	yAxisLength := len(grid)

	for _, slope := range slopes {
		trees := 0
		x, y := 0, 0
		for range grid {
			y += slope.down
			x += slope.right
			if x > xAxisLength-1 {
				x -= xAxisLength
			}
			if string(grid[y][x]) == "#" {
				trees++
			}
			if yAxisLength-1 == y {
				break
			}
		}
		answer *= trees
	}

	return strconv.Itoa(answer)

}
