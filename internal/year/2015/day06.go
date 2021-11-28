package aoc2015

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

const (
	ON     = "on"
	OFF    = "off"
	TOGGLE = "toggle"
)

type (
	Coords struct {
		X, Y int
	}
	Command struct {
		Action   string
		From, To Coords
	}
	LitLights map[Coords]int

	day06 struct{}
)

func Day06() solution.Provider {
	return &day06{}
}

func (*day06) SolveA(input string) string {
	ll := &LitLights{}
	commands := strings.Split(input, "\n")
	for _, c := range commands {
		if c == "" {
			continue
		}
		ll = execCommand(*ll, parseCommand(c))
	}

	return strconv.Itoa(len(*ll))
}

func parseCommand(cmdLine string) Command {
	c := Command{}
	part := strings.Split(cmdLine, " ")
	switch part[0] {
	case TOGGLE:
		c.Action = TOGGLE
		c.From = parseCoords(part[1])
		c.To = parseCoords(part[3])
	case "turn":
		if c.Action = ON; part[1] == OFF {
			c.Action = OFF
		}
		c.From = parseCoords(part[2])
		c.To = parseCoords(part[4])
	}

	return c
}

func parseCoords(coords string) Coords {
	s := strings.Split(coords, ",")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])

	return Coords{x, y}
}

func execCommand(ll LitLights, c Command) *LitLights {
	switch c.Action {
	case ON:
		for y := c.From.Y; y <= c.To.Y; y++ {
			for x := c.From.X; x <= c.To.X; x++ {
				ll[Coords{x, y}] = 1
			}
		}
	case OFF:
		for y := c.From.Y; y <= c.To.Y; y++ {
			for x := c.From.X; x <= c.To.X; x++ {
				delete(ll, Coords{x, y})
			}
		}
	case TOGGLE:
		for y := c.From.Y; y <= c.To.Y; y++ {
			for x := c.From.X; x <= c.To.X; x++ {
				c := Coords{x, y}
				if _, ok := ll[c]; ok {
					delete(ll, c)
				} else {
					ll[c] = 1
				}
			}
		}
	}

	return &ll
}

func (*day06) SolveB(input string) string {
	var answer int
	ll := &LitLights{}
	cmds := strings.Split(input, "\n")
	for _, cmd := range cmds {
		if cmd == "" {
			continue
		}
		ll = execCommandB(*ll, parseCommand(cmd))
	}

	for _, v := range *ll {
		answer += v
	}

	return strconv.Itoa(answer)
}

func execCommandB(ll LitLights, cmd Command) *LitLights {
	switch cmd.Action {
	case "on":
		for y := cmd.From.Y; y <= cmd.To.Y; y++ {
			for x := cmd.From.X; x <= cmd.To.X; x++ {
				ll[Coords{x, y}]++
			}
		}
	case "off":
		for y := cmd.From.Y; y <= cmd.To.Y; y++ {
			for x := cmd.From.X; x <= cmd.To.X; x++ {
				c := Coords{x, y}
				if ll[c] > 0 {
					ll[c]--
				}
			}
		}
	case "toggle":
		for y := cmd.From.Y; y <= cmd.To.Y; y++ {
			for x := cmd.From.X; x <= cmd.To.X; x++ {
				ll[Coords{x, y}] += 2
			}
		}
	}

	return &ll
}
