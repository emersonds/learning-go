// https://leetcode.com/problems/fizz-buzz/submissions/1936690141

package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	str := []string{}

	// Loop i to n, checking if it is divisible by 3, 5, or both
	for i := 1; i <= n; i++ {
		// If divisible by 3 and 5, this index should be FizzBuzz
		if (i % 3 == 0) && (i % 5 == 0) {
			str = append(str, "FizzBuzz")
		} else if (i % 3 == 0) {
			str = append(str, "Fizz")
		} else if (i % 5 == 0) {
			str = append(str, "Buzz")
		} else {
			str = append(str, strconv.Itoa(i))
		}
	}

	fmt.Println(str)
	return str
}

func main() {
	// Test cases
	fizzBuzz(3)
	fizzBuzz(5)
	fizzBuzz(15)
}
