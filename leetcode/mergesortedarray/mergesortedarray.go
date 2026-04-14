package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// Used to start nums2 at index 0
	nIndex := 0

	// Loop through array, starting at m and ending at m+n
	// m is how many elements to be merged (nums1),
	// n is elements that are set to 0 and should be ignored (change 0s to nums2 values)
	for i := m; i < m+n; i++ {
		// Start at first 0, set it to nums2 value
		nums1[i] = nums2[nIndex]
		// Increment nums2 index
		nIndex++
	}

	sort.Ints(nums1)

	return nums1
}

func main() {
	testNums1 := []int{0}
	testNums2 := []int{1}
	m := 0
	n := 1

	fmt.Println(merge(testNums1, m, testNums2, n))
}
