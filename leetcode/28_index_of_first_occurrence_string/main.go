// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/submissions/1987936509
package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}
