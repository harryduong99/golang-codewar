package main

import "log"

const V = "victory"
const D = "defeat"

var (
	inQueue []int64
	subs    []int64
)

func Solution(total int64, conditions [][]int64) string {
	// Write your solution here
	setSubs(conditions)
	var prequisites []int64
	var result bool
	for _, missions := range conditions {
		prequisites = append(prequisites, missions[1])
		result = recursiveChecking(conditions, missions[1])
		inQueue = []int64{}
		if !result {
			return D
		}
	}

	return V
}

func setSubs(conditions [][]int64) {
	for _, missions := range conditions {
		subs = append(subs, missions[0])
	}
}

func recursiveChecking(conditions [][]int64, toAccomplish int64) bool {
	for _, missions := range conditions {
		if toAccomplish == missions[0] {
			if contains(inQueue, toAccomplish) {
				return false
			}
			inQueue = append(inQueue, toAccomplish)
			return recursiveChecking(conditions, missions[1])
		}
	}

	return true
}

func contains(s []int64, target int64) bool {
	for _, a := range s {
		if a == target {
			return true
		}
	}
	return false
}

func main() {
	log.Println(Solution(3, [][]int64{{1, 2}, {2, 3}, {3, 5}}))
}
