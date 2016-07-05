package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{
		4,
		&ListNode{
			5,
			nil,
		},
	}
	b := &ListNode{
		5,
		nil,
	}
	printListNode(a)
	printListNode(b)
	printListNode(addTwoNumbers(a, b))
}

func printListNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	result := &ListNode{}
	pre := result
	head := result
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			result.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			result.Val += l2.Val
			l2 = l2.Next
		}
		result.Val += carry
		carry = result.Val / 10
		result.Val = result.Val % 10
		result.Next = &ListNode{}
		pre = result
		result = result.Next
	}
	pre.Next = nil
	return head
}
