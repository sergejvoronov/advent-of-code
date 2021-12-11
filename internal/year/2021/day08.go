package aoc2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"

	"github.com/juliangruber/go-intersect"
)

type day08 struct{}

func Day08() solution.Provider {
	return &day08{}
}

func (*day08) SolveA(input string) string {
	var (
		answer int
		digits = map[int]int{
			2: 1,
			4: 4,
			3: 7,
			7: 8,
		}
	)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		content := strings.Split(line, " | ")
		output := strings.Fields(content[1])
		for _, n := range output {
			if _, ok := digits[len(n)]; ok {
				answer++
			}
		}
	}
	return strconv.Itoa(answer)
}

func (*day08) SolveB(input string) string {

	//      0:      1:      2:      3:      4:
	//     aaaa    ....    aaaa    aaaa    ....
	//    b    c  .    c  .    c  .    c  b    c
	//    b    c  .    c  .    c  .    c  b    c
	//     ....    ....    dddd    dddd    dddd
	//    e    f  .    f  e    .  .    f  .    f
	//    e    f  .    f  e    .  .    f  .    f
	//     gggg    ....    gggg    gggg    ....

	//      5:      6:      7:      8:      9:
	//     aaaa    aaaa    aaaa    aaaa    aaaa
	//    b    .  b    .  .    c  b    c  b    c
	//    b    .  b    .  .    c  b    c  b    c
	//     dddd    dddd    ....    dddd    dddd
	//    .    f  e    f  .    f  e    f  .    f
	//    .    f  e    f  .    f  e    f  .    f
	//     gggg    gggg    ....    gggg    gggg

	// Unique digits based on their number of segments:  1, 4, 7, and 8
	// We can deduce the other digits based on how many segments they share
	// with the unique digits (excluding 8 bc it has all segments)

	// +-------+-----+----+----+----+
	// | digit | len | ∩1 | ∩7 | ∩4 |
	// +-------+-----+----+----+----+
	// |     0 |   6 |  2 |  3 |  3 |
	// |     2 |   5 |  1 |  2 |  2 |
	// |     3 |   5 |  2 |  3 |  3 |
	// |     5 |   5 |  1 |  2 |  3 |
	// |     6 |   6 |  1 |  2 |  3 |
	// |     9 |   6 |  2 |  3 |  4 |
	// +-------+-----+----+----+----+

	var (
		content []string
		answer  int
		digits  = make(map[int]string)
	)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		content = strings.Split(line, " | ")
		patterns := strings.Fields(content[0])
		for _, p := range patterns {
			switch len(p) {
			case 2:
				digits[1] = p
			case 3:
				digits[7] = p
			case 4:
				digits[4] = p
			case 7:
				digits[8] = p
			}
		}

		for _, p := range patterns {
			switch len(p) {
			case 5:
				if len(intersect.Simple(p, digits[4])) == 2 {
					digits[2] = p
					continue
				}
				if len(intersect.Simple(p, digits[1])) == 2 {
					digits[3] = p
					continue
				}
				digits[5] = p
			case 6:
				if len(intersect.Simple(p, digits[1])) == 1 {
					digits[6] = p
					continue
				}
				if len(intersect.Simple(p, digits[4])) == 4 {
					digits[9] = p
					continue
				}
				digits[0] = p
			}
		}

		res := ""
		decodings := strings.Fields(content[1])

		for _, decoding := range decodings {
			for k, v := range digits {
				// Letters are not ordered, so need to check for each
				commonLetters := intersect.Simple(strings.Split(v, ""), strings.Split(decoding, ""))
				if len(commonLetters) == len(v) && len(commonLetters) == len(decoding) {
					res += fmt.Sprintf("%d", k)
				}
			}
		}

		num, _ := strconv.Atoi(res)
		answer += num
	}

	return strconv.Itoa(answer)
}
