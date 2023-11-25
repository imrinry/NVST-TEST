package main

func findOddInt(arr []int) int {
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}

	for num, count := range counts {
		if count%2 != 0 {
			return num
		}
	}

	return -1
}
