package main

import (
	"fmt"
	"strings"
)

func FirstNonRepeating(str string) string {
	toLower := strings.ToLower(str)
	originChars := strings.Split(str, "")
	chars := strings.Split(toLower, "")
	var count int
	var toRemove []int
	for i, c := range chars {
		if _, isExisting := Find(toRemove, i); isExisting {
			continue
		}

		count = 0
		for j, rc := range chars[i+1:] {
			if rc == c {
				toRemove = append(toRemove, i+j+1)
			} else {
				count++
			}
		}

		if count == len(chars)-i-1 {
			return originChars[i]
		}
	}
	return ""
}

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func main() {
	fmt.Println(FirstNonRepeating("sTreSS"))
}
