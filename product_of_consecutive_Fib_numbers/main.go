package main

import (
	"fmt"
)

func ProductFib(prod uint64) [3]uint64 {
	// your code
	var result [3]uint64
	var fib, fibB, fibC uint64
	var i uint64 = 0
	for {
		i++
		fib = Fibo(i)
		fibB = Fibo(i + 1)
		fibC = Fibo(i + 2)
		if fib*fibB == prod {
			result[0] = fib
			result[1] = fibB
			result[2] = 1
			return result
		} else if fib*fibB < prod && fibB*fibC > prod {
			result[0] = fibB
			result[1] = fibC
			result[2] = 0
			return result
		}
	}

}

func Fibo(n uint64) uint64 {
	if n == 0 || n == 1 {
		return n
	}

	return Fibo(n-1) + Fibo(n-2)
}

func FindNearestFibo(x uint64) uint64 {
	var i uint64 = 0
	var fib uint64
	for {
		i++
		fib = Fibo(i)
		if fib > x {
			return Fibo(i - 1)
		}
	}
}

func main() {
	// fmt.Println(Fibo(8))
	fmt.Println(ProductFib(800))
}
