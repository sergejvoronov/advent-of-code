package aoc2015

import (
	"strconv"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day03 struct{}

func Day03() solution.Provider {
	return &day03{}
}

func (*day03) SolveA(input string) string {
	type House struct {
		X, Y int
	}

	h := House{0, 0}
	presents := make(map[House]int)
	presents[h]++
	for _, direction := range input {
		switch string(direction) {
		case ">":
			h.X++
		case "<":
			h.X--
		case "^":
			h.Y++
		case "v":
			h.Y--
		default:
			continue
		}
		presents[h]++
	}

	return strconv.Itoa(len(presents))
}

func (*day03) SolveB(input string) string {
	type (
		House   struct{ X, Y int }
		Carrier struct{ House }
	)
	var c *Carrier

	presents := make(map[Carrier]int)
	santa := Carrier{House{0, 0}}
	robot := Carrier{House{0, 0}}
	presents[santa]++
	presents[robot]++
	for nextMove, direction := range input {
		if c = &robot; nextMove%2 == 0 {
			c = &santa
		}
		switch string(direction) {
		case ">":
			c.House.X++
		case "<":
			c.House.X--
		case "^":
			c.House.Y++
		case "v":
			c.House.Y--
		default:
			continue
		}
		presents[*c]++
	}

	return strconv.Itoa(len(presents))
}
