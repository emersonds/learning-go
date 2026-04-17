// https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/1981318272
package main

import (
	"fmt"
)

func removeDuplicates(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, []int{}
	}

	k := 1

	// Check each element until there is no duplicates
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k, nums
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
