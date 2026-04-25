// https://leetcode.com/problems/linked-list-cycle/submissions/1987920586
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// Given the head of a linked list, determine if linked list has a cycle

	// Return true if there is a cycle, otherwise false

	// Example, head = [3,2,0,-4], pos = 1, where pos is the index of the node
	// that the tail node's "next" pointer is connected to. pos is not passed
	// as a parameter.

	// save reference to head
	var first *ListNode = head
	
	// make hashtable (map) of *ListNodes to check for cycles against
	// The way this works is that the key is a reference to the first appearance
	// of a ListNode. If the map contains the ListNode, we have seen it before,
	// so the value is true. This is because the zero-value of bools is false.
	var nodes = make(map[*ListNode]bool)

	// Loop through each ListNode
	for i := first; i != nil; i = i.Next {
		// if this ListNode has already appeared, the value is true, so return true
		if nodes[i] {
			return true
		}
		// the current ListNode has not been seen, so we add it to the map.
		nodes[i] = true
	}
	
	// there is no cycle, so return false
	return false
}
