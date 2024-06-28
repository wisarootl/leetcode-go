package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/move-zeroes/description/

func moveZeroes(nums []int) []int {
	lastNonZeroFoundAt := 0

	// Move all the non-zero elements forward
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp := nums[lastNonZeroFoundAt]
			nums[lastNonZeroFoundAt] = nums[i]
			nums[i] = temp

			lastNonZeroFoundAt++
		}
	}

	return nums

}

func test(nums []int, expected []int) {
	output := moveZeroes(nums)
	if reflect.DeepEqual(output, expected) {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func main() {
	test([]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0})
	test([]int{0}, []int{0})
}
