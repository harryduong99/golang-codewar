package main

import "log"

var (
	resultA []int
	resultJ []int
	totalA  int
	totalJ  int
	totalsA []int
	totalsJ []int
)

func Ann(n int) []int {
	var result []int
	var t int
	var toAppend int
	for i := 0; i < n; i++ {
		if i == 0 {
			toAppend = 1
		} else {
			t = result[i-1]
			toAppend = i - John(t + 1)[t]
		}
		result = append(result, toAppend)

		if len(resultA)-1 < i {
			resultA = append(resultA, toAppend)
			totalA += toAppend
			totalsA = append(totalsA, totalA)
		}
	}

	return result
}

func John(n int) []int {
	var result []int
	var t int
	var toAppend int
	for i := 0; i < n; i++ {
		if i == 0 {
			toAppend = 0
		} else {
			t = resultJ[i-1]
			toAppend = i - Ann(t + 1)[t]
		}

		result = append(result, toAppend)
		if len(resultJ)-1 < i {
			resultJ = append(resultJ, toAppend)
			totalJ += toAppend
			totalsJ = append(totalsJ, totalJ)
		}
	}

	return result
}

func SumJohn(n int) int {
	John(n)
	return totalsJ[n-1]
}

func SumAnn(n int) int {
	Ann(n)
	return totalsA[n-1]
}

func main() {
	log.Println(Ann(6))
	log.Println(SumAnn(6))
}
