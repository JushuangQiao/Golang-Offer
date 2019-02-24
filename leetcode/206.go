package leetcode

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre = nil
	for head != nil {
		next := head.Next
		head.Next = pre
		pre, head = head, next
	}
	return pre
}
