package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-n]
	}
	fmt.Println(nums1)
}

func main() {
	testNums1 := []int{1, 2, 3, 0, 0, 0}
	testNums2 := []int{2, 5, 6}
}
