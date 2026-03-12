package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left := 0
	right := n - 1
	pos := n - 1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[pos] = leftSquare
			left++
		} else {
			result[pos] = rightSquare
			right--
		}

		pos--
	}

	return result
}

func main() {
	nums1 := []int{-4, -1, 0, 3, 10}

	nums1 = sortedSquares(nums1)

	fmt.Println(nums1)
}
