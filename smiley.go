package main

import (
	"fmt"
	"regexp"
)

func countSmileys(arr []string) int {
	count := 0
	for _, face := range arr {
		match, _ := regexp.MatchString(`^[:;][-~]?[)D]`, face)
		if match {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(countSmileys([]string{":)", ";(", ";}", ":-D"}))
	fmt.Println(countSmileys([]string{";D", ":-(", ":-)", ";~)"}))
	fmt.Println(countSmileys([]string{";]", ":[", ";*", ":$", ";-D"}))
}
