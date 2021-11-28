package app

import (
	"flag"
	"log"

	"github.com/sergejvoronov/advent-of-code/internal/config"
	"github.com/sergejvoronov/advent-of-code/internal/input"
	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

var (
	yearValue = flag.String("year", "", "Advent of Code year value")
	dayValue  = flag.String("day", "", "Advent of Code day value of the year (will not work without year)")
)

func Exec() error {
	flag.Parse()

	c, err := config.New("config.yaml")
	if err != nil {
		log.Fatalf("Config file failed to load: %v", err)
	}

	inputReader := input.NewReader(c.App.Session)
	solver := solution.NewSolver(inputReader, config.SolutionsByYear)

	return solver.Solve(*yearValue, *dayValue)
}
