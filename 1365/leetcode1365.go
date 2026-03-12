package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	freq := make([]int, 101)
	result := make([]int, n)

	for _, num := range nums {
		freq[num]++
	}

	for i := 1; i < 101; i++ {
		freq[i] += freq[i-1]
	}

	for i, num := range nums {
		if num == 0 {
			result[i] = 0
		} else {
			result[i] = freq[num-1]
		}
	}

	return result
}

func main() {
	nums := []int{8, 1, 2, 2, 3}

	nums = smallerNumbersThanCurrent(nums)

	fmt.Println(nums)
}
