package main

import (
	"fmt"	
)

var steps int = 0
var isCounting bool = false

func numberOfSteps(num int) int {
	// Number is greater than 0 and first time executing function
	if !isCounting {
		isCounting = true
		steps = 0
	}
	
	fmt.Println("Entered numberOfSteps:", isCounting, steps, num)
	
	// Don't go below zero
	if (num > 0) {
		// If a number is even, divide by two
		if (num % 2 == 0) {
			num = num / 2
			steps++
			// Recurse until 0
			numberOfSteps(num)
		// If number is odd, subtract one
		} else if (num % 2 != 0 && num != 0) {
			num--
			steps++
			// Recurse until 0
			numberOfSteps(num)
		}
	}

	isCounting = false	// Executed after recursion to reset steps
	
	return steps
}

func main() {
	fmt.Println(numberOfSteps(14))	// Expected 6
	fmt.Println(numberOfSteps(8))		// Expected 4
	fmt.Println(numberOfSteps(0))
}
