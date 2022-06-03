package main

import (
	"log"
	"strings"
)

func isBalanced(s string) string {
	// Write your code here
	brackets := strings.Split(s, "")
	var stack []string
	var top string

	for _, c := range brackets {
		if c == ")" || c == "]" || c == "}" {
			if len(stack) == 0 {
				return "NO"
			}
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (c == ")" && top != "(") || (c == "]" && top != "[") || (c == "}" && top != "{") {
				return "NO"
			}
		} else {
			stack = append(stack, c)
		}
	}

	return "YES"
}

func main() {
	demo := "}([[{)[]))]{){}["
	log.Println(isBalanced(demo))
}
