package aoc2021

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day07 struct{}

func Day07() solution.Provider {
	return &day07{}
}

func (*day07) SolveA(input string) string {
	content := strings.Split(input, ",")

	var (
		requiredFuel, median int
		positions            = make([]int, 0, len(content))
	)

	// Convert to ints
	for _, pos := range content {
		posInt, _ := strconv.Atoi(pos)
		positions = append(positions, posInt)
	}

	// Identify most frequently used position (median)
	sort.Ints(positions)
	middle := len(positions) / 2
	if len(positions)%2 != 0 {
		median = positions[middle]
	} else {
		median = (positions[middle-1] + positions[middle]) / 2
	}

	// Calculate fuel consumption
	for _, p := range positions {
		if p == median {
			continue
		}
		requiredFuel += int(math.Abs(float64(p - median)))
	}

	return strconv.Itoa(requiredFuel)
}

func (*day07) SolveB(input string) string {
	// input = `16,1,2,0,4,2,7,1,2,14`

	content := strings.Split(input, ",")

	var (
		requiredFuel, sum, mean int
		positions               = make([]int, 0, len(content))
	)

	// Convert to ints
	for _, pos := range content {
		posInt, _ := strconv.Atoi(pos)
		positions = append(positions, posInt)
	}

	// Calculate mean position
	for _, pos := range positions {
		sum += pos
	}
	mean = int(math.Floor(float64(sum) / float64(len(positions))))

	// Arithmetic progression function
	calculateFuel := func(n int) int {
		return (n * (1 + n)) / 2
	}

	// Calculate minimal fuel consumption
	for _, p := range positions {
		if p == mean {
			continue
		}
		distance := int(math.Abs(float64(p - mean)))
		requiredFuel += calculateFuel(distance)
	}

	return strconv.Itoa(requiredFuel)
}
