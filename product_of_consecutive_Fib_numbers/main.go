package main

import (
	"fmt"
	"math"
)

var fibos []int

func ProductFib(prod uint64) [3]uint64 {
	// your code
	aver := uint64(math.Sqrt(float64(prod)))
	var result [3]uint64
	var fib, fibStore uint64 = 0, 0
	var i uint64 = 0
	for {
		i++
		fib = Fibonacci(i)
		if fib*fibStore == prod {
			result[0] = fibStore
			result[1] = fib
			result[2] = 1
			return result
		}
		if fib > aver {
			if fib*fibStore > prod {
				result[0] = fibStore
				result[1] = fib
			} else {
				result[0] = fib
				result[1] = Fibonacci(i + 1)
			}
			result[2] = 0
			return result
		}
		fibStore = fib
	}
}

func Fibo(n uint64) uint64 {
	if n == 0 || n == 1 {
		return n
	}

	return Fibo(n-1) + Fibo(n-2)
}

func Fibonacci(n uint64) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	var n2, n1 uint64 = 0, 1

	for i := uint64(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
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

// FINEST PRACTICE!!!
func ProductFibBest(prod uint64) [3]uint64 {
	f1 := uint64(0)
	f2 := uint64(1)
	for f1*f2 < prod {
		f1, f2 = f2, f1+f2
	}
	success := uint64(0)
	if prod == f1*f2 {
		success = 1
	}
	return [3]uint64{f1, f2, success}
}
