package aoc2021

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day14 struct{}

func Day14() solution.Provider {
	return &day14{}
}

func (*day14) SolveA(input string) string {
	var (
		answers     = make([]int, 0)
		rules       = make(map[string]string)
		occurrences = make(map[rune]int)
		template    string
	)

	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			template = line
			continue
		}
		if i == 1 {
			continue
		}
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	for step := 0; step < 10; step++ {
		stepTemplate := ""
		for i := 0; i < len(template)-1; i++ {
			if element, ok := rules[template[i:i+2]]; ok {
				stepTemplate += string(template[i]) + element
				if i == len(template)-2 {
					stepTemplate += string(template[len(template)-1])
				}

			}
		}
		template = stepTemplate
	}

	for _, e := range template {
		occurrences[e]++
	}

	for _, o := range occurrences {
		answers = append(answers, o)
	}

	sort.Ints(answers)

	return strconv.Itoa(answers[len(answers)-1] - answers[0])
}

func (*day14) SolveB(input string) string {
	var (
		rules    = make(map[string]string)
		template string
	)

	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			template = line
			continue
		}
		if i == 1 {
			continue
		}
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for i := 0; i < 40; i++ {
		newPairs := map[string]int{}
		for pair := range pairs {
			newPairs[string(pair[0])+rules[pair]] += pairs[pair]
			newPairs[rules[pair]+string(pair[1])] += pairs[pair]
		}
		pairs = newPairs
	}

	counter := make(map[string]int)

	for i := range pairs {
		counter[string(i[0])] += pairs[i]
	}
	counter[string(template[len(template)-1])] += 1

	min := math.MaxInt
	max := 0

	for _, v := range counter {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return strconv.Itoa(max - min)
}
