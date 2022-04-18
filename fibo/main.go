package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	f1 := 0
	f2 := 1
	i := 2
	for i <= n {
		i++
		f1, f2 = f2, f1+f2
	}

	return f2
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacci(n))
}
