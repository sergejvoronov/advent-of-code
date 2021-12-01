package aoc2021

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

	for k := range numbers {
		if k == 0 {
			continue
		}
		if numbers[k] > numbers[k-1] {
			answer++
		}
	}

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

	numberLength := len(numbers)

	for k := range numbers {
		if k+3 == numberLength {
			break
		}
		if numbers[k+1]+numbers[k+2]+numbers[k+3] > numbers[k]+numbers[k+1]+numbers[k+2] {
			answer++
		}
	}

	return strconv.Itoa(answer)
}
