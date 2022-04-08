package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 2}
	find(test)
}

func find(nums []int) []int {
	result := []int{}
	fmt.Println(getDupNum(nums))
	return result
}

func getDupNum(nums []int) int {
	for _, i := range nums {
		for _, j := range nums[i+1:] {
			if i == j {
				return i
			}
		}
	}
	return 0
}

func gettotal(n int) int {
	return (n + 1) * n / 2
}

func getTotalll(nums []int) int {
	total := 0
	for _, i := range nums {
		total += i
	}

	return total
}
