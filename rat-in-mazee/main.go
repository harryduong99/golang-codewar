package main

import (
	"fmt"
	"strconv"
)

const N int = 4

var (
	sol  [N][N]int
	maze [N][N]int
)

func printSolution(sol [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf(strconv.Itoa(sol[i][j]))
			fmt.Printf(" ")
		}
		fmt.Println("")
	}
}

func isSafe(sol [N][N]int, x int, y int) bool {
	if x >= 0 && x < N && y >= 0 && y < N && sol[x][y] == 1 {
		return true
	}

	return false
}

func solveMaze(maze [N][N]int) bool {
	// fmt.Println(solveMazeUtil(maze, 0, 0, sol))
	if solveMazeUtil(maze, 0, 0, sol) == false {
		return false
	}

	return true
}

func solveMazeUtil(maze [N][N]int, x int, y int, sol [N][N]int) bool {
	// if (x, y is goal) return true
	if x == N-1 && y == N-1 && maze[x][y] == 1 {
		sol[x][y] = 1
		printSolution(sol)
		return true
	}
	// Check if maze[x][y] is valid
	if isSafe(maze, x, y) {
		// Check if the current block is already part of solution path.
		if sol[x][y] == 1 {
			return false
		}

		sol[x][y] = 1

		/* Move forward in x direction */
		if solveMazeUtil(maze, x+1, y, sol) {
			return true
		}

		/* If moving in x direction doesn't give
		solution then Move down in y direction */
		if solveMazeUtil(maze, x, y+1, sol) {
			return true
		}

		/* If moving in y direction doesn't give
		solution then Move backwards in x direction */
		if solveMazeUtil(maze, x-1, y, sol) {
			return true
		}

		/* If moving backwards in x direction doesn't give
		solution then Move upwards in y direction */
		if solveMazeUtil(maze, x, y-1, sol) {
			return true
		}

		/* If none of the above movements works then
		   BACKTRACK: unmark x, y as part of solution
		   path */
		sol[x][y] = 0
		return false
	}
	return false
}

func initSol() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sol[i][j] = 0
		}
	}
}

func main() {
	maze = [N][N]int{{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1}}

	initSol()
	solveMaze(maze)
}
