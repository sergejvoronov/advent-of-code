package aoc2021

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day11 struct{}

func Day11() solution.Provider {
	return &day11{}
}

func (*day11) SolveA(input string) string {
	var (
		flashes   int
		steps     int = 100
		octopuses     = make([][]int, 0)
	)

	lines := strings.Split(input, "\n")
	for y := range lines {
		var line []int

		for _, s := range lines[y] {
			n, _ := strconv.Atoi(string(s))
			line = append(line, n)
		}
		octopuses = append(octopuses, line)
	}

	for s := 0; s < steps; s++ {
		for y := range octopuses {
			for x := range octopuses[y] {
				flashes += checkPoint(octopuses, 0, y, x)
			}
		}
		for y := range octopuses {
			for x := range octopuses[y] {
				if octopuses[y][x] == -1 {
					octopuses[y][x] = 0
				}
			}
		}
	}

	return strconv.Itoa(flashes)
}

func (*day11) SolveB(input string) string {
	var (
		allFlashes int = 100
		flashes    int
		answer     int
		octopuses  = make([][]int, 0)
	)

	lines := strings.Split(input, "\n")
	for y := range lines {
		var line []int

		for _, s := range lines[y] {
			n, _ := strconv.Atoi(string(s))
			line = append(line, n)
		}
		octopuses = append(octopuses, line)
	}

	step := 0
	for flashes != allFlashes {
		flashes = 0
		step++
		for y := range octopuses {
			for x := range octopuses[y] {
				flashes += checkPoint(octopuses, 0, y, x)
			}
		}
		for y := range octopuses {
			for x := range octopuses[y] {
				if octopuses[y][x] == -1 {
					octopuses[y][x] = 0
				}
			}
		}
		answer = step
	}

	return strconv.Itoa(answer)
}

func checkPoint(octopuses [][]int, flashes, y, x int) int {
	if octopuses[y][x] == -1 {
		return flashes
	}

	if octopuses[y][x] != 9 {
		octopuses[y][x]++
		return flashes
	}

	octopuses[y][x] = -1
	flashes++
	if y > 0 {
		if x > 0 {
			flashes = checkPoint(octopuses, flashes, y-1, x-1)
		}
		flashes = checkPoint(octopuses, flashes, y-1, x)
		if x < len(octopuses[y-1])-1 {
			flashes = checkPoint(octopuses, flashes, y-1, x+1)
		}
	}

	if x > 0 {
		flashes = checkPoint(octopuses, flashes, y, x-1)
	}
	if x < len(octopuses[y])-1 {
		flashes = checkPoint(octopuses, flashes, y, x+1)
	}

	if y < len(octopuses)-1 {
		if x > 0 {
			flashes = checkPoint(octopuses, flashes, y+1, x-1)
		}
		flashes = checkPoint(octopuses, flashes, y+1, x)
		if x < len(octopuses[y+1])-1 {
			flashes = checkPoint(octopuses, flashes, y+1, x+1)
		}
	}

	return flashes
}
