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
func isSafe() bool {
	return false
}

func solve(col int) bool {
	for i := 0; i < N; i++ {
		if isSafe() {
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
