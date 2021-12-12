package aoc2021

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day10 struct{}

func Day10() solution.Provider {
	return &day10{}
}

func (*day10) SolveA(input string) string {
	var (
		answer     int
		appearance = make(map[string]int)
		scores     = map[string]int{
			")": 3,
			"]": 57,
			"}": 1197,
			">": 25137,
		}
	)

	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`(\(\)|\[\]|\<\>|\{\})`)
	invalid := regexp.MustCompile(`(\)|\]|\>|\})`)
	for i := range lines {
		for re.MatchString(lines[i]) {
			lines[i] = re.ReplaceAllString(lines[i], "")
		}

		symbol := invalid.FindString(lines[i])
		if symbol != "" {
			appearance[symbol]++
		}
	}

	for symbol, appeared := range appearance {
		answer += appeared * scores[symbol]
	}

	return strconv.Itoa(answer)
}

func (*day10) SolveB(input string) string {
	var (
		answers = make([]int, 0)
		answer  int
		scores  = map[string]int{
			"(": 1,
			"[": 2,
			"{": 3,
			"<": 4,
		}
	)

	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`(\(\)|\[\]|\<\>|\{\})`)
	invalid := regexp.MustCompile(`(\)|\]|\>|\})`)
	for i := range lines {
		for re.MatchString(lines[i]) {
			lines[i] = re.ReplaceAllString(lines[i], "")
		}

		if invalid.FindString(lines[i]) != "" {
			continue
		}

		answer = 0
		for x := len(lines[i]) - 1; x >= 0; x-- {
			answer = 5*answer + scores[string(lines[i][x])]
		}
		answers = append(answers, answer)
	}

	sort.Ints(answers)

	return strconv.Itoa(answers[len(answers)/2])
}
