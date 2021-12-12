package aoc2021

import (
	"sort"
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day09 struct{}

func Day09() solution.Provider {
	return &day09{}
}

func (*day09) SolveA(input string) string {
	var answer int

	lines := strings.Split(input, "\n")
	matrix := make(map[int][]int)

	for y, line := range lines {
		for _, val := range line {
			intVal, _ := strconv.Atoi(string(val))
			matrix[y] = append(matrix[y], intVal)
		}
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if x > 0 {
				if matrix[y][x-1] <= matrix[y][x] {
					continue
				}
			}
			if x < len(matrix[y])-1 {
				if matrix[y][x+1] <= matrix[y][x] {
					continue
				}
			}
			if y > 0 {
				if matrix[y-1][x] <= matrix[y][x] {
					continue
				}
			}
			if y < len(lines)-1 {
				if matrix[y+1][x] <= matrix[y][x] {
					continue
				}
			}

			answer += matrix[y][x] + 1
		}
	}

	return strconv.Itoa(answer)
}

func (*day09) SolveB(input string) string {
	var basins []int

	lines := strings.Split(input, "\n")
	matrix := make(map[int][]int)

	for y, line := range lines {
		for _, val := range line {
			intVal, _ := strconv.Atoi(string(val))
			matrix[y] = append(matrix[y], intVal)
		}
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if x > 0 {
				if matrix[y][x-1] <= matrix[y][x] {
					continue
				}
			}
			if x < len(matrix[y])-1 {
				if matrix[y][x+1] <= matrix[y][x] {
					continue
				}
			}
			if y > 0 {
				if matrix[y-1][x] <= matrix[y][x] {
					continue
				}
			}
			if y < len(lines)-1 {
				if matrix[y+1][x] <= matrix[y][x] {
					continue
				}
			}

			_, count := exploreBasins(matrix, 0, y, x)

			basins = append(basins, count)
		}
	}

	sort.Ints(basins)

	return strconv.Itoa(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}

func exploreBasins(matrix map[int][]int, count, y, x int) (map[int][]int, int) {
	if matrix[y][x] != 9 {
		count++
		matrix[y][x] = 9
		if y != len(matrix)-1 {
			matrix, count = exploreBasins(matrix, count, y+1, x)
		}
		if y != 0 {
			matrix, count = exploreBasins(matrix, count, y-1, x)
		}
		if x != len(matrix[y])-1 {
			matrix, count = exploreBasins(matrix, count, y, x+1)
		}
		if x != 0 {
			matrix, count = exploreBasins(matrix, count, y, x-1)
		}
	}

	return matrix, count
}
