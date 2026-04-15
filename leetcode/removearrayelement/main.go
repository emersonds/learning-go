package main

import "fmt"

// Removes all occurrences of val in nums
// Returns number of elements in nums that are not equal to val
func removeElement(nums []int, val int) int {
	k := 0
	n := len(nums)

	for k < n {
		if nums[k] == val {
			// Swap val element with last element
			nums[k] = nums[n-1]
			// Last element is val now, go to previous element
			n--
		} else {
			k++
		}
	}

	return k
}

func main() {
	testNums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	testVal1 := 2
	fmt.Println(removeElement(testNums1, testVal1))
}
