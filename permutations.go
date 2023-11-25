package main

import (
	"sort"
)

func generatePermutations(input string) []string {
	var result []string
	permutate([]rune(input), 0, &result)
	result = removeDuplicates(result)
	sort.Strings(result)
	return result
}

func permutate(s []rune, index int, result *[]string) {
	if index == len(s)-1 {
		*result = append(*result, string(s))
		return
	}

	for i := index; i < len(s); i++ {
		s[index], s[i] = s[i], s[index]
		permutate(s, index+1, result)
		s[index], s[i] = s[i], s[index]
	}
}

func removeDuplicates(arr []string) []string {
	encountered := map[string]struct{}{}
	result := []string{}

	for _, val := range arr {
		if _, ok := encountered[val]; !ok {
			encountered[val] = struct{}{}
			result = append(result, val)
		}
	}

	return result
}
