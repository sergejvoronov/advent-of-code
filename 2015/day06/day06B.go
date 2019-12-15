package day06

import (
	"fmt"
	"strings"
)

func SolveB(input string) string {
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

	return fmt.Sprintf("Day 06(B) answer: %d", answer)
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
