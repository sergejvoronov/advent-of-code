package solution

import (
	"fmt"
	"strconv"

	"github.com/sergejvoronov/advent-of-code/internal/input"
)

type (
	Solver interface {
		Solve(year, day string) error
	}

	solver struct {
		inputReader       input.Reader
		solutionProviders map[string][]Provider
	}
)

func NewSolver(
	inputReader input.Reader,
	solutionProviders map[string][]Provider,

) Solver {
	return &solver{
		inputReader:       inputReader,
		solutionProviders: solutionProviders,
	}
}

func (s *solver) Solve(year, day string) error {
	if year == "" {
		for year, solutions := range s.solutionProviders {
			_ = s.showSolutionsByDay(solutions, year, day)
		}

		return nil
	}

	if _, ok := s.solutionProviders[year]; !ok {
		return fmt.Errorf("Solutions not found for provided year: %v", year)
	}
	solutions := s.solutionProviders[year]

	return s.showSolutionsByDay(solutions, year, day)

}

func (s *solver) showSolutionsByDay(solutions []Provider, year, day string) error {
	fmt.Printf("Year %v solutions:\n", year)
	if day != "" {
		dayInt, err := strconv.Atoi(day)
		if err != nil && day != "" {
			return fmt.Errorf("Provided day value is invalid: %v", day)
		}
		if dayInt == 0 || dayInt > len(solutions) {
			return fmt.Errorf("Solution not found for provided day: %v", day)
		}
		s.showDaySolutions(year, dayInt, solutions[dayInt-1])

		return nil
	}

	for day, solution := range solutions {
		s.showDaySolutions(year, day+1, solution)
	}

	return nil
}

func (s *solver) showDaySolutions(year string, day int, provider Provider) {
	dayString := strconv.Itoa(day)
	dayInput, err := s.inputReader.ReadInput(year, dayString)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	fmt.Printf("Day %s answer (A): %s\n", dayString, provider.SolveA(dayInput))
	fmt.Printf("Day %s answer (B): %s\n", dayString, provider.SolveB(dayInput))
}
