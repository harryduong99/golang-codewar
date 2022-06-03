package main

import "log"

func lastStoneWeight(weights []int32) int32 {
	// Write your code here
	// var stonesLeft []int32
	if len(weights) == 0 {
		return 0
	}
	if len(weights) == 2 && containsInt32(weights, 1) {
		largestElement, _, _ := largestElement(weights)

		return largestElement
	}
	_, positions, secondLargestPosition := largestElement(weights)

	if len(positions) >= 2 {
		weights = initNewWeights(weights, positions)
		return lastStoneWeight(weights)
	}
	left := weights[positions[0]] - weights[secondLargestPosition]

	positions = append(positions, secondLargestPosition)
	weights = initNewWeights(weights, positions)
	weights = append(weights, left)

	if len(weights) == 1 {
		return weights[0]
	}

	return lastStoneWeight(weights)
}

func initNewWeights(weights []int32, positions []int) []int32 {
	var newWeight []int32
	for i, v := range weights {
		if !contains(positions, i) {
			newWeight = append(newWeight, v)
		}
	}

	return newWeight
}

func contains(s []int, target int) bool {
	for _, a := range s {
		if a == target {
			return true
		}
	}
	return false
}

func containsInt32(s []int32, target int32) bool {
	for _, a := range s {
		if a == target {
			return true
		}
	}
	return false
}

func largestElement(array []int32) (int32, []int, int) {
	max := array[0]
	maxPosition := 0
	var secondLargestPosition int
	var secondLargest int32
	var positions []int
	positions = append(positions, 0)
	if len(array) == 1 {
		return 0, []int{0}, 0
	}

	for i := 1; i < len(array); i++ {
		if max < array[i] {
			secondLargestPosition = maxPosition
			secondLargest = max
			maxPosition = i
			max = array[i]
			positions = []int{}
			positions = append(positions, i)
			continue
		} else if secondLargest < array[i] {
			secondLargestPosition = i
			secondLargest = array[i]
		}

		if max == array[i] {
			positions = append(positions, i)
		}
	}

	return max, positions, secondLargestPosition
}

func main() {
	arr := []int32{1, 974961061}
	// arr := []int32{32, 46188086, 339992587, 742976890, 801915058, 393898202, 717833291, 843435009, 361066046, 884145908, 668431192, 586679703, 792103686, 85425451, 246993674, 134274127, 586374055, 923288873, 292845117, 399188845, 842456591, 410257930, 333998862, 16561419, 624279391, 459765367, 969764608, 508221973, 82956997, 437034793, 553121267, 554066040, 199416087}
	log.Println(lastStoneWeight(arr))

}
