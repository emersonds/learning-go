// https://leetcode.com/submissions/detail/1943379578/?from=explore&item_id=3240

package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	for index, value := range nums {
		nums[index] = value * value
	}

	sort.Ints(nums)

	return nums
}

func main() {
	test1 := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(test1))
}
