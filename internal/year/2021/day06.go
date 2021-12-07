package aoc2021

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day06 struct{}

func Day06() solution.Provider {
	return &day06{}
}

func (d *day06) SolveA(input string) string {
	const cycle = 80

	return strconv.Itoa(d.spawnFishes(input, cycle))
}

func (d *day06) SolveB(input string) string {
	const cycle = 256

	return strconv.Itoa(d.spawnFishes(input, cycle))
}

func (d *day06) spawnFishes(input string, cycle int) (fishCount int) {
	timers := strings.Split(input, ",")
	fishes := make([]int, 0)
	for _, t := range timers {
		timer, _ := strconv.Atoi(t)
		fishes = append(fishes, timer)
	}

	days := make([]int, 9)
	for _, fish := range fishes {
		days[fish]++
	}
	for i := 0; i < cycle; i++ {
		activeFish := days[0]
		for i := 0; i < 8; i++ {
			days[i] = days[i+1]
		}
		days[6] += activeFish
		days[8] = activeFish
	}

	for _, fish := range days {
		fishCount += fish
	}

	return fishCount
}
