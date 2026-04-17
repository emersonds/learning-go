package main

import "fmt"

func checkIfExist(arr []int) bool {
	for i := range arr {
		for j := 1; j < len(arr); j++ {
			if arr[i] == 2 * arr[j-1] {
				fmt.Println(arr[i], arr[j], 2 * arr[j])
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkIfExist([]int{0,-2,2}))
}
