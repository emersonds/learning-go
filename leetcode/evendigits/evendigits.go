package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		var newStr string = strconv.Itoa(nums[i])
		if len(newStr) % 2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	test1 := []int{555, 901, 482, 1771}
	test2 := []int{5164, 920, 3205, 32}
	fmt.Println(findNumbers(test1))  // 1
	fmt.Println(findNumbers(test2))  // 3
}
