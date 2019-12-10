package day03

import (
	"fmt"
)

func SolveB(input string) string {
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

	return fmt.Sprintf("Day 03(B) answer: %d", len(presents))
}
