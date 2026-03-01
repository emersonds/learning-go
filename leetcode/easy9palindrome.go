package main

import (
	//"fmt"
	"strconv"
)

func isPalinedrome(x int) bool {
	// negative ints are not palindromes: -131 is 131-
	if x < 0 {
		return false
	}
	
	// Convert int to string
	strInt := strconv.Itoa(x)
	//fmt.Println("strInt:", strInt)

	// Create backwards string to compare
	backwards := []rune(strInt)
	// i = 0, j = length of string
	// Loop while i < j, increment i and decrement j each loop
	for i, j := 0, len(backwards)-1; i < j; i, j = i+1, j-1 {
		backwards[i], backwards[j] = backwards[j], backwards[i]
	}

	//fmt.Println("backwards:", string(backwards))

	// Compare new string to backwards string
	return strInt == string(backwards)
}

func main () {
	isPalinedrome(121)
	isPalinedrome(149)
	isPalinedrome(-131)
}

// https://leetcode.com/problems/palindrome-number/submissions/1935064107
