package main

import (
  "fmt"
)

func Pyramid(n int) [][]int {
    // your code here
	a := make([][]int, n)
	for i := n; i > 0; i-- {
		var b []int
		for j := i; j > 0; j-- {
			b = append(b, 1)
		} 
		a[i-1] = b
		
	}
	
	return a
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	Pyramid(3)
}
