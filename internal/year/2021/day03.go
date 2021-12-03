package aoc2021

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day03 struct{}

type day03Counter struct {
	zeros, ones             int
	zeroNumbers, oneNumbers []string
}

func Day03() solution.Provider {
	return &day03{}
}

func (d *day03) SolveA(input string) string {
	var gamma, epsilon string

	numbers := strings.Split(input, "\n")

	for bitNumber := 0; bitNumber < len(numbers[0]); bitNumber++ {
		c := d.countZerosAndOnes(numbers, bitNumber)

		if c.ones > c.zeros {
			gamma += "0"
			epsilon += "1"
			continue
		}
		gamma += "1"
		epsilon += "0"

	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	return strconv.Itoa(int(gammaInt) * int(epsilonInt))
}

func (d *day03) SolveB(input string) string {
	var oxygenGenerator, co2scrubber string

	numbers := strings.Split(input, "\n")

	oxygenFilteredNumbers, co2scrubberFilteredNumbers := numbers, numbers

	for bitNumber := 0; bitNumber < len(numbers[0]); bitNumber++ {
		c := d.countZerosAndOnes(oxygenFilteredNumbers, bitNumber)

		if c.ones > c.zeros || c.ones == c.zeros {
			oxygenFilteredNumbers = c.oneNumbers
		} else {
			oxygenFilteredNumbers = c.zeroNumbers
		}

		if len(oxygenFilteredNumbers) == 1 {
			oxygenGenerator = oxygenFilteredNumbers[0]
			break
		}
	}

	for bitNumber := 0; bitNumber < len(numbers[0]); bitNumber++ {
		c := d.countZerosAndOnes(co2scrubberFilteredNumbers, bitNumber)

		if c.ones > c.zeros || c.ones == c.zeros {
			co2scrubberFilteredNumbers = c.zeroNumbers
		} else {
			co2scrubberFilteredNumbers = c.oneNumbers
		}

		if len(co2scrubberFilteredNumbers) == 1 {
			co2scrubber = co2scrubberFilteredNumbers[0]
			break
		}
	}

	oxygenGeneratorInt, _ := strconv.ParseInt(oxygenGenerator, 2, 64)
	co2scrubberInt, _ := strconv.ParseInt(co2scrubber, 2, 64)

	return strconv.Itoa(int(oxygenGeneratorInt) * int(co2scrubberInt))
}

func (*day03) countZerosAndOnes(numbers []string, bit int) day03Counter {
	c := day03Counter{}

	for _, number := range numbers {
		if string(number[bit]) == "0" {
			c.zeros++
			c.zeroNumbers = append(c.zeroNumbers, number)
			continue
		}
		c.ones++
		c.oneNumbers = append(c.oneNumbers, number)
	}

	return c
}
