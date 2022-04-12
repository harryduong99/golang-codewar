package main

import "log"

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
			t = result[i-1]
			toAppend = i - Ann(t + 1)[t]
		}
		result = append(result, toAppend)
	}

	return result
}

func SumJohn(n int) int {
	result := John(n)
	totalJ := 0
	for _, i := range result {
		totalJ += i
	}

	return totalJ
}

func SumAnn(n int) int {
	result := Ann(n)
	totalA := 0
	for _, i := range result {
		totalA += i
	}

	return totalA
}

func main() {
	log.Println(Ann(6))
	log.Println(SumAnn(6))
}
