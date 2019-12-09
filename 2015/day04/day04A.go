package day04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func SolveA(input string) string {
	var answer int
	input = strings.TrimSpace(input)
	for {
		str := strconv.Itoa(answer)
		hash := md5.Sum([]byte(input + str))
		if strings.Compare(hex.EncodeToString(hash[:])[:5], "00000") == 0 {
			break
		}
		answer++
	}

	return fmt.Sprintf("Day 04(A) answer: %d", answer)
}
