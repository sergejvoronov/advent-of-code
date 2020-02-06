package day08

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func SolveA(input string) string {
	charCount, inMemory := 0, 0
	xs := strings.Split(input, "\n")
	for _, s := range xs {
		if s == "" {
			continue
		}
		charCount += len(s)
		inMemory += utf8.RuneCountInString(s)
	}
	return fmt.Sprintf("Day 08(A) answer: %d", charCount-inMemory)
}
