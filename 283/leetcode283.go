package main

import "fmt"

func moveZeroes(nums []int) {
	k := 0

	for i, num := range nums {
		if num != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 3, 12}

	moveZeroes(nums)

	fmt.Println(nums)
}

