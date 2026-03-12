package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)

	sumExpected := (n * (n + 1)) / 2

	sumArr := 0

	for _, v := range nums {
		sumArr += v
	}

	return sumExpected - sumArr
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	result := missingNumber(nums)

	fmt.Println(result)
}
