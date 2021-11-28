package aoc2015

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day07 struct{}

func Day07() solution.Provider {
	return &day07{}
}

func (*day07) SolveA(input string) string {
	wires := make(map[string]int)
	signals := strings.Split(input, "\n")
	counter := 0
	for counter != len(signals) {
		for _, signal := range signals[counter:] {
			counter++
			if signal == "" {
				continue
			}
			s := strings.Split(signal, " -> ")
			cmd := strings.Split(s[0], " ")
			switch len(cmd) {
			case 1:
				if v, err := strconv.Atoi(cmd[0]); err == nil {
					wires[s[1]] = v
				} else if _, ok := wires[cmd[0]]; !ok {
					postponeSignal(&signals, signal)
				} else {
					wires[s[1]] = wires[cmd[0]]
				}
			case 2:
				if _, ok := wires[cmd[1]]; !ok {
					postponeSignal(&signals, signal)
					continue
				}
				wires[s[1]] = int(^uint16(wires[cmd[1]]))
			case 3:
				switch cmd[1] {
				case "AND":
					var firstNum int
					if v, err := strconv.Atoi(cmd[0]); err == nil {
						firstNum = v
					} else if _, ok := wires[cmd[0]]; !ok {
						postponeSignal(&signals, signal)
						continue
					} else {
						firstNum = wires[cmd[0]]
					}
					if _, ok := wires[cmd[2]]; !ok {
						postponeSignal(&signals, signal)
						continue
					}
					wires[s[1]] = firstNum & wires[cmd[2]]
				case "OR":
					_, firstExists := wires[cmd[0]]
					_, secondExists := wires[cmd[2]]
					if !firstExists || !secondExists {
						postponeSignal(&signals, signal)
						continue
					}
					wires[s[1]] = wires[cmd[0]] | wires[cmd[2]]
				case "LSHIFT":
					if _, ok := wires[cmd[0]]; !ok {
						postponeSignal(&signals, signal)
						continue
					}
					v, _ := strconv.Atoi(cmd[2])
					wires[s[1]] = wires[cmd[0]] << v
				case "RSHIFT":
					if _, ok := wires[cmd[0]]; !ok {
						postponeSignal(&signals, signal)
						continue
					}
					v, _ := strconv.Atoi(cmd[2])
					wires[s[1]] = wires[cmd[0]] >> v
				}
			default:
				panic(fmt.Sprintf("Command '%s' cannot be parsed", s[0]))
			}
		}
	}
	return strconv.Itoa(wires["a"])
}

func postponeSignal(signals *[]string, signal string) {
	*signals = append(*signals, signal)
}

func (*day07) SolveB(input string) string {
	wires := make(map[string]int)
	signals := strings.Split(input, "\n")
	secondStep := false
	counter := 0
	for counter != len(signals) {
		for _, signal := range signals[counter:] {
			counter++
			if signal == "" {
				continue
			}
			s := strings.Split(signal, " -> ")
			if s[1] == "b" && secondStep {
				continue
			}
			cmd := strings.Split(s[0], " ")
			switch len(cmd) {
			case 1:
				if v, err := strconv.Atoi(cmd[0]); err == nil {
					wires[s[1]] = v
				} else if _, ok := wires[cmd[0]]; !ok {
					postponeSignal(&signals, signal)
				} else {
					wires[s[1]] = wires[cmd[0]]
				}
			case 2:
				if _, ok := wires[cmd[1]]; !ok {
					postponeSignal(&signals, signal)
					continue
				}
				wires[s[1]] = int(^uint16(wires[cmd[1]]))
			case 3:
				switch cmd[1] {
				case "AND":
					var firstNum int
					if v, err := strconv.Atoi(cmd[0]); err == nil {
						firstNum = v
					} else if _, ok := wires[cmd[0]]; !ok {
						postponeSignal(&signals, signal)
						continue
					} else {
						firstNum = wires[cmd[0]]
					}
					if _, ok := wires[cmd[2]]; !ok {
						postponeSignal(&signals, signal)
						continue
					}
					wires[s[1]] = firstNum & wires[cmd[2]]
				case "OR":
					_, firstExists := wires[cmd[0]]
					_, secondExists := wires[cmd[2]]
					if !firstExists || !secondExists {
						postponeSignal(&signals, signal)
						continue
					}
					wires[s[1]] = wires[cmd[0]] | wires[cmd[2]]
				case "LSHIFT":
					if _, ok := wires[cmd[0]]; !ok {
						postponeSignal(&signals, signal)
						continue
					}
					v, _ := strconv.Atoi(cmd[2])
					wires[s[1]] = wires[cmd[0]] << v
				case "RSHIFT":
					if _, ok := wires[cmd[0]]; !ok {
						postponeSignal(&signals, signal)
						continue
					}
					v, _ := strconv.Atoi(cmd[2])
					wires[s[1]] = wires[cmd[0]] >> v
				}
			default:
				panic(fmt.Sprintf("Command '%s' cannot be parsed", s[0]))
			}
		}
		if v, ok := wires["a"]; ok {
			if !secondStep {
				secondStep = true
				for k := range wires {
					delete(wires, k)
				}
				wires["b"] = v
				signals = strings.Split(input, "\n")
				counter = 0
				continue
			} else {
				break
			}
		}
	}
	return strconv.Itoa(wires["a"])
}
