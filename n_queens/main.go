package main

import (
	"fmt"
	"strconv"
)

const N int = 4

var board [N][N]int

func initBoard() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			board[i][j] = 0
		}
	}
}

func printBoard() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf(strconv.Itoa(board[i][j]))
			fmt.Printf(" ")
		}
		fmt.Println("")
	}
}

// check if the queen won't be check
func isSafe(row int, col int) bool {
	for i := 0; i <= N; i++ {
		if board[row][i] == 1 {
			return false
		}
	}

	for r, c := row, col; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == 1 {
			return false
		}
	}

	for a, b := row, col; b >= 0 && a < N; a, b = a+1, b-1 {
		if board[a][b] == 1 {
			return false
		}
	}

	return true

}

func solve(col int) bool {
	for i := 0; i < N; i++ {
		if isSafe(i, col) {
			board[i][col] = 1
			if solve(col + 1) {
				return true
			}

			board[i][col] = 0
		}
	}

	return false
}

func solveQueen() bool {
	if solve(0) {
		printBoard()
		return true
	}

	return false
}

func main() {
	initBoard()
	solveQueen()
}
