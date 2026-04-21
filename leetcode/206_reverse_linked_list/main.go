// https://leetcode.com/problems/reverse-linked-list/submissions/1984868033
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// Need three pointers to reverse,
	// pointers to prev node, curr node, and next node
	var prev *ListNode = nil
	var curr *ListNode = head
	var next *ListNode

	// iterate through nodes until curr node is nil
	for curr != nil {
		// store next node for updating curr to next pointer
		next = curr.Next
		// Set actual next node to previous node, reversing the link
		curr.Next = prev
		// set pointer to prev node to the current node
		// and set pointer to curr node to next node
		// to iterate to next node in list
		prev = curr
		curr = next
	}

	// Return prev node because curr node is nil
	// and prev node starts the reversed list
	return prev
}
