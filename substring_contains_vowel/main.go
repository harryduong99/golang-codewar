package main

import (
	"log"
	"strings"
)

var (
	vowel []string
	total int = 0
)

func vowelsubstring(s string) int32 {
	// Write your code here
	var i, j int
	c := strings.Split(s, "")
	n := len(s)
	for i = 0; i < n; i++ {
		for j = i; j < n; j++ {
			if !isVowel(c[j]) {
				break
			}

			if !contains(vowel, c[j]) {
				vowel = append(vowel, c[j])
			}

			if len(vowel) == 5 {
				total++
			}
		}
		vowel = []string{}
	}

	return int32(total)
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func isVowel(x string) bool {
	if x == "a" || x == "e" || x == "i" || x == "o" || x == "u" {
		return true
	}

	return false

}

func main() {
	s := "aaeiouxa"
	log.Println(vowelsubstring(s))
}
