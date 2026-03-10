// https://leetcode.com/submissions/detail/1943391734/?from=explore&item_id=3245

package main

import (
	"fmt"
)

func duplicateZeros(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			// Move each element to the right
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			// Skip the next index because the next index will be 0
			i++
		}
	}

	return arr
}

func main() {
	test1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	test2 := []int{1, 2, 3}

	fmt.Println(duplicateZeros(test1))
	fmt.Println(duplicateZeros(test2))
}
