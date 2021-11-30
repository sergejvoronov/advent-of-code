package config

import (
	"github.com/sergejvoronov/advent-of-code/internal/solution"
	ac2015 "github.com/sergejvoronov/advent-of-code/internal/year/2015"
	ac2020 "github.com/sergejvoronov/advent-of-code/internal/year/2020"
	ac2021 "github.com/sergejvoronov/advent-of-code/internal/year/2021"
)

var (
	SolutionsByYear = map[string][]solution.Provider{
		"2021": ac2021.Solutions,
		"2020": ac2020.Solutions,
		"2015": ac2015.Solutions,
	}
)
