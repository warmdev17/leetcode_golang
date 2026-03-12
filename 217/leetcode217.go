package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, v := range nums {
		if seen[v] {
			return true
		}
		seen[v] = true
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}

	result := containsDuplicate(nums)

	fmt.Println(result)
}
