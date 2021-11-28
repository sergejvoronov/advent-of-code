package aoc2015

import (
	"crypto/md5" //nolint
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

type day04 struct{}

func Day04() solution.Provider {
	return &day04{}
}

func (*day04) SolveA(input string) string {
	var answer int
	input = strings.TrimSpace(input)
	for {
		str := strconv.Itoa(answer)
		hash := md5.Sum([]byte(input + str)) //nolint
		if strings.Compare(hex.EncodeToString(hash[:])[:5], "00000") == 0 {
			break
		}
		answer++
	}

	return strconv.Itoa(answer)
}

func (*day04) SolveB(input string) string {
	var answer int
	input = strings.TrimSpace(input)
	for {
		str := strconv.Itoa(answer)
		hash := md5.Sum([]byte(input + str)) //nolint
		if strings.Compare(hex.EncodeToString(hash[:])[:6], "000000") == 0 {
			break
		}
		answer++
	}

	return strconv.Itoa(answer)
}
