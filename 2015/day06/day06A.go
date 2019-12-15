package day06

import (
	"fmt"
	"strconv"
	"strings"
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
)

func SolveA(input string) string {
	ll := &LitLights{}
	cmds := strings.Split(input, "\n")
	for _, cmd := range cmds {
		if cmd == "" {
			continue
		}
		ll = execCommand(*ll, parseCommand(cmd))
	}

	return fmt.Sprintf("Day 06(A) answer: %d", len(*ll))
}

func parseCommand(cmdLine string) Command {
	cmd := Command{}
	part := strings.Split(cmdLine, " ")
	switch part[0] {
	case "toggle":
		cmd.Action = "toggle"
		cmd.From = parseCoords(part[1])
		cmd.To = parseCoords(part[3])
	case "turn":
		if cmd.Action = "on"; part[1] == "off" {
			cmd.Action = "off"
		}
		cmd.From = parseCoords(part[2])
		cmd.To = parseCoords(part[4])
	}

	return cmd
}

func parseCoords(coords string) Coords {
	s := strings.Split(coords, ",")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])

	return Coords{x, y}
}

func execCommand(ll LitLights, cmd Command) *LitLights {
	switch cmd.Action {
	case "on":
		for y := cmd.From.Y; y <= cmd.To.Y; y++ {
			for x := cmd.From.X; x <= cmd.To.X; x++ {
				ll[Coords{x, y}] = 1
			}
		}
	case "off":
		for y := cmd.From.Y; y <= cmd.To.Y; y++ {
			for x := cmd.From.X; x <= cmd.To.X; x++ {
				delete(ll, Coords{x, y})
			}
		}
	case "toggle":
		for y := cmd.From.Y; y <= cmd.To.Y; y++ {
			for x := cmd.From.X; x <= cmd.To.X; x++ {
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
