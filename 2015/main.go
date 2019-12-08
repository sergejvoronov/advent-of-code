package main

import (
	"github.com/sergejvoronov/adventofcode"
	"github.com/sergejvoronov/adventofcode/2015/day01"
	"github.com/sergejvoronov/adventofcode/2015/day02"
)

func main() {
	// DAY 01
	println(day01.SolveA(adventofcode.ReadInputData("inputs/day01.input")))
	println(day01.SolveB(adventofcode.ReadInputData("inputs/day01.input")))
	// DAY 02
	println(day02.SolveA(adventofcode.ReadInputData("inputs/day02.input")))
	println(day02.SolveB(adventofcode.ReadInputData("inputs/day02.input")))
}
