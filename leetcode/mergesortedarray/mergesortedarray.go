package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	for i := m; i < m+n; i++ {
		// Compiler error if m = 0
		nums1[i] = nums2[i-n]
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
