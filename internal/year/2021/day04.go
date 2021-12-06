package aoc2021

import (
	"strconv"
	"strings"

	"github.com/sergejvoronov/advent-of-code/internal/solution"
)

var wonBoardKeys = make(map[int]struct{})

type day04 struct{}

func Day04() solution.Provider {
	return &day04{}
}

func (d *day04) SolveA(input string) string {
	boards, bingoNumbers := d.parseInputToBoardsAndBingoNumbers(input)
	winningBoard, winningBingoNumber := d.getFirstWinningBoard(boards, bingoNumbers)
	unmarkedNumbersSum := sumUnmarkedNumbers(winningBoard)

	return strconv.Itoa(unmarkedNumbersSum * winningBingoNumber)
}

func (d *day04) SolveB(input string) string {
	boards, bingoNumbers := d.parseInputToBoardsAndBingoNumbers(input)
	lastWinningBoard, lastWinningBingoNumber := d.getLastWinningBoard(boards, bingoNumbers)
	unmarkedNumbersSum := sumUnmarkedNumbers(lastWinningBoard)

	return strconv.Itoa(unmarkedNumbersSum * lastWinningBingoNumber)
}

func (*day04) parseInputToBoardsAndBingoNumbers(input string) ([][][]int, []int) {
	var (
		boards = make([][][]int, 0)
		board  = make([][]int, 0)
	)

	content := strings.Split(input, "\n")
	bingoSequence := content[0]

	for i, line := range content {
		if i < 2 {
			continue
		}
		if line == "" {
			boards = append(boards, board)
			board = make([][]int, 0)
			continue
		}
		lineNumbers := strings.Fields(line)
		lineNumbersInt := make([]int, 0)
		for _, lineNumber := range lineNumbers {
			lineNumberInt, _ := strconv.Atoi(lineNumber)
			lineNumbersInt = append(lineNumbersInt, lineNumberInt)
		}
		board = append(board, lineNumbersInt)
	}
	// Add last board
	boards = append(boards, board)

	// Parse bingo sequence numbers
	sequenceNumbers := make([]int, 0)
	for _, sequenceNumber := range strings.Split(bingoSequence, ",") {
		sequenceNumberInt, _ := strconv.Atoi(sequenceNumber)
		sequenceNumbers = append(sequenceNumbers, sequenceNumberInt)
	}

	return boards, sequenceNumbers
}

func (d *day04) getFirstWinningBoard(boards [][][]int, bingoNumbers []int) ([][]int, int) {
	for _, bingoNumber := range bingoNumbers {
		for boardKey, board := range boards {
			for row, line := range board {
				for column := range line {
					if board[row][column] == bingoNumber {
						board[row][column] = -1
						// Get first winning board
						if board, won := d.checkWinningBoard(board, boardKey, row, column); won {
							return board, bingoNumber
						}
					}
				}
			}
		}
	}

	return [][]int{}, 0
}

func (d *day04) getLastWinningBoard(boards [][][]int, bingoNumbers []int) ([][]int, int) {
	boardsLeft := len(boards)
	for _, bingoNumber := range bingoNumbers {
		for boardKey, board := range boards {
			for row, line := range board {
				for column := range line {
					if board[row][column] == bingoNumber {
						board[row][column] = -1
						// Get last winning board
						if _, won := d.checkWinningBoard(board, boardKey, row, column); won {
							boardsLeft--
							if boardsLeft == 1 {
								return board, bingoNumber
							}
						}
					}
				}
			}
		}
	}

	return [][]int{}, 0
}

func (*day04) checkWinningBoard(board [][]int, boardKey int, row, column int) ([][]int, bool) {
	if _, alreadyWon := wonBoardKeys[boardKey]; alreadyWon {
		return nil, false
	}
	// Check vertical lines
	winningColumn := func(board [][]int) bool {
		for x := 0; x < 5; x++ {
			if board[x][column] != -1 {
				return false
			}
		}
		return true
	}(board)

	// Check horizontal lines
	winningRow := func(board [][]int) bool {
		for x := 0; x < 5; x++ {
			if board[row][x] != -1 {
				return false
			}
		}
		return true
	}(board)

	if !winningColumn && !winningRow {
		return nil, false
	}

	wonBoardKeys[boardKey] = struct{}{}
	return board, true
}

func sumUnmarkedNumbers(board [][]int) (unmarkedNumbersSum int) {
	for row, line := range board {
		for column := range line {
			if board[row][column] != -1 {
				unmarkedNumbersSum += board[row][column]
			}
		}
	}

	return unmarkedNumbersSum
}
