package main

import (
	"fmt"
	"math"
)

// func ProductFib(prod uint64) [3]uint64 {
// 	// your code
// 	var result [3]uint64
// 	var fib, fibB, fibC uint64
// 	var i uint64 = 0
// 	for {
// 		i++
// 		fib = Fibo(i)
// 		fibB = Fibo(i + 1)
// 		fibC = Fibo(i + 2)
// 		if fib*fibB == prod {
// 			result[0] = fib
// 			result[1] = fibB
// 			result[2] = 1
// 			return result
// 		} else if fib*fibB < prod && fibB*fibC > prod {
// 			result[0] = fibB
// 			result[1] = fibC
// 			result[2] = 0
// 			return result
// 		}
// 	}

// }

func ProductFib(prod uint64) [3]uint64 {
	// your code
	aver := uint64(math.Sqrt(float64(prod)))
	var result [3]uint64
	var fib, fibStore uint64 = 0, 0
	var i uint64 = 0
	for {
		i++
		fib = Fibo(i)
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
				result[1] = Fibo(i + 1)
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
