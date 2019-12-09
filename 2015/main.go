package main

import (
	"github.com/sergejvoronov/adventofcode"
	"github.com/sergejvoronov/adventofcode/2015/day01"
	"github.com/sergejvoronov/adventofcode/2015/day02"
	"github.com/sergejvoronov/adventofcode/2015/day03"
	"github.com/sergejvoronov/adventofcode/2015/day04"
)

func main() {
	// DAY 04
	println(day04.SolveA(adventofcode.ReadInputData("inputs/day04.input")))
	println(day04.SolveB(adventofcode.ReadInputData("inputs/day04.input")))
	// DAY 03
	println(day03.SolveA(adventofcode.ReadInputData("inputs/day03.input")))
	println(day03.SolveB(adventofcode.ReadInputData("inputs/day03.input")))
	// DAY 02
	println(day02.SolveA(adventofcode.ReadInputData("inputs/day02.input")))
	println(day02.SolveB(adventofcode.ReadInputData("inputs/day02.input")))
	// DAY 01
	println(day01.SolveA(adventofcode.ReadInputData("inputs/day01.input")))
	println(day01.SolveB(adventofcode.ReadInputData("inputs/day01.input")))
}
