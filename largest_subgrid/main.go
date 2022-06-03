package main

import (
	"log"
	"math"
)

func largestSubgrid(grid [][]int32, maxSum int32) int32 {
	// Write your code here
	var i, j int = 0, 0
	x := len(grid)
	y := len(grid[0])

	sum := preparePrefixSum(x, y)
	sum = updatePrefixSum(grid, sum, x, y)
	var result int = 0

	for i = 1; i <= x; i++ {
		for j = 1; j <= y; j++ {
			if i+result-1 > x || j+result-1 > y {
				break
			}

			low := result
			hi := min(x-i+1, y-j+1)

			for low < hi {
				mid := int(math.Floor(float64(hi+low+1)) / 2)
				if sum[i+mid-1][j+mid-1]+sum[i-1][j-1]-sum[i+mid-1][j-1]-sum[i-1][j+mid-1] <= maxSum {
					low = mid
				} else {
					hi = mid - 1
				}
			}
			result = max(result, low)
		}
	}

	return int32(result)
}

func preparePrefixSum(x int, y int) [][]int32 {
	var sum [][]int32

	for i := 0; i < x+1; i++ {
		sum = append(sum, []int32{})
		for j := 0; j < y+1; j++ {
			sum[i] = append(sum[i], 0)

		}
	}

	return sum
}

func updatePrefixSum(grid [][]int32, sum [][]int32, x int, y int) [][]int32 {
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			if i == 0 || j == 0 {
				sum[i][j] = 0
				continue
			}

			sum[i][j] = grid[i-1][j-1] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := [][]int32{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	log.Println(largestSubgrid(a, 4))
}
