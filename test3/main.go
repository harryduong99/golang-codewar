package main

import "log"

func Solution(array []int64) []int64 {
	var current int64
	var result []int64
	// Write your solution here
	for i, v := range array {
		block := []int64{}
		if i == 0 {
			current = v
			if array[i+1] <= current {
				result = append(result, int64(i))
			}
			continue
		}

		if v <= current && v <= array[i-1] {
			block = append(block, v)
			continue
		} else {
			current = v
			if isLastItem(array, i) {
				result = append(result, int64(i))
				break
			}

			// for j, k := range array[i:] {

			// }
			if array[i+1] <= current {
				result = append(result, int64(i))
			}
		}
	}

	return result
}

func isLastItem(array []int64, i int) bool {
	return i+1 == len(array)
}

func main() {
	// a := []int64{10, 5, 2, 3, 11, 7}
	// a := []int64{3, 7, 2, 3, 7, 11, 15, 8, 11, 9}
	// a := []int64{5, 93, 35, 69, 69, 56, 35, 24, 85, 89, 70, 88, 70, 10, 55}
	// a := []int64{20, 67, 41, 7, 5, 92, 60, 52, 2, 36}
	a := []int64{83, 70, 95, 24, 47, 26, 15, 38, 18, 24, 84, 84, 92, 75, 7}
	log.Println(Solution(a))
}
