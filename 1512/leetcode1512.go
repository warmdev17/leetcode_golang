package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	pairCount := 0
	freq := make(map[int]int)

	for _, num := range nums {
		pairCount += freq[num]
		freq[num]++
	}

	return pairCount
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}

	result := numIdenticalPairs(nums)

	fmt.Println(result)
}
