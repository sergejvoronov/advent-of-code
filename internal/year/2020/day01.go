package aoc2020

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day01 struct{}

func Day01() solution.Provider {
	return &day01{}
}

func (*day01) SolveA(input string) string {
	var answer int

	strings := strings.Split(input, "\n")
	numbers := make([]int, 0, len(strings))
	for _, s := range strings {
		if s == "" {
			break
		}
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	func() {
		for x := range numbers {
			for y := range numbers {
				if y <= x {
					continue
				}
				if numbers[x]+numbers[y] == 2020 {
					answer = numbers[x] * numbers[y]

					return
				}
			}
		}
	}()

	return strconv.Itoa(answer)
}

func (*day01) SolveB(input string) string {
	var answer int

	strings := strings.Split(input, "\n")
	numbers := make([]int, 0, len(strings))
	for _, s := range strings {
		if s == "" {
			break
		}
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	func() {
		for x := range numbers {
			for y := range numbers {
				if y <= x {
					continue
				}
				for z := range numbers {
					if z <= y {
						continue
					}
					if numbers[x]+numbers[y]+numbers[z] == 2020 {
						answer = numbers[x] * numbers[y] * numbers[z]

						return
					}
				}
			}
		}
	}()

	return strconv.Itoa(answer)
}
